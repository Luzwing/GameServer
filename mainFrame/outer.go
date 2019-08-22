package mainFrame

import (
	"fmt"
	"goproto/appclose"
	"goproto/gameEnter"
	"goproto/roomMessage"
	"io"
	"math/rand"
	"messageType"
	"net"
	"notice"
	"playerstatus"
	"runtime"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

func Process(conn net.Conn) error {
	fmt.Println("进入房间外处理进程")
	for {
		//读取消息长度
		var msgLenByte []byte
		//msgLenByte := make([]byte, 4)
		var lenLenHaveRead int32 = 0
		for {
			tempBuf := make([]byte, 4-lenLenHaveRead)
			lenByte, e := conn.Read(tempBuf[0 : 4-lenLenHaveRead])

			if e != nil {
				if e == io.EOF {
					//runtime.Goexit()
					break
				}
				Slog.Log2file(e.Error())
				fmt.Println("Receive Failed, err:", e)
			}
			lenLenHaveRead += int32(lenByte)
			msgLenByte = CombineBytes(msgLenByte, tempBuf)
			if lenLenHaveRead == 4 {
				break
			}
		}
		//fmt.Println(msgLenByte)

		msglen := BytesToInt(msgLenByte)
		fmt.Println("消息长度：", msglen)

		var buf []byte
		//buf := make([]byte, msglen)
		//已经读到的字节数
		var protoLenHaveRead int32 = 0
		for {
			tempBuf := make([]byte, msglen-protoLenHaveRead)
			protoLen, err := conn.Read(tempBuf[0:(msglen - protoLenHaveRead)])

			if err != nil {
				if err == io.EOF {
					//runtime.Goexit()
					break
				}
				Slog.Log2file(err.Error())
				fmt.Println("Receive Failed, err:", err)
			}
			protoLenHaveRead += int32(protoLen)
			buf = CombineBytes(buf, tempBuf)
			if protoLenHaveRead == msglen {
				break
			}
		}

		msgType := buf[0]
		msgContent := buf[1:msglen]

		switch msgType {
		//玩家登录
		case messageType.C_GAMEENTER:
			playerEnterGame(msgContent, conn)
		//房间列表获取
		case messageType.C_ROOMGET:
			getAllRoom(msgContent, conn)
		//房间创建
		case messageType.C_ROOMCRT:
			crtRoom(msgContent, conn)
		//房间加入
		case messageType.C_ROOMJOIN:
			joinRoom(msgContent, conn)
		//有客户端强制退出
		case messageType.C_APPCLOSE:
			appCloseOuterRoom(msgContent)
		//快速加入
		case messageType.C_QUICKJOIN:
			quickJoin(msgContent, conn)
		default:
			var flush []byte
			conn.Read(flush)
		}
	}

	return nil
}

func playerEnterGame(msgContent []byte, conn net.Conn) {
	msgUnmarshalled := &gameEnter.C2GS_GameEnter{}
	proto.Unmarshal(msgContent, msgUnmarshalled)

	//若该玩家ID已存在且该玩家没有在房间中，则先从队列中删除，后续加锁
	PlayersArrMutex.Lock()
	defer PlayersArrMutex.Unlock()
	// for index, player := range playersArr {
	// 	if (player.playerId == msgUnmarshalled.GetPlayerId()) && (player.playerName == msgUnmarshalled.GetPlayerName()) {
	// 		// if player.status == playerstatus.NotInRoom {
	// 		// 	playersArr = append(playersArr[:index], playersArr[index+1:]...)
	// 		// 	player.conn.Close()
	// 		// }
	// 		//加锁
	// 		playersArr = append(playersArr[:index], playersArr[index+1:]...)
	// 		//player.conn.Close()
	// 	}
	// }

	newPlayer := Player{
		conn:       conn,
		playerId:   msgUnmarshalled.GetPlayerId(),
		playerName: msgUnmarshalled.GetPlayerName(),
		status:     playerstatus.NotInRoom,
		roomId:     0,
	}
	playersArr = append(playersArr, newPlayer)

	msg2Marshal := &gameEnter.GS2C_GameEnter{}
	msg2Marshal.ResStatus = 100
	msgContent2Send, _ := proto.Marshal(msg2Marshal)
	msg2Send := TC_Combine(messageType.S_GAMEENTER, msgContent2Send)
	_, e := sendMsg2One(conn, msg2Send)
	if e != nil {
		Slog.Log2file("login send error")
	}
	fmt.Printf("目前还剩%d位玩家在线\n", len(playersArr))
	for i, p := range playersArr {
		fmt.Printf("第%d位用户\n:", i+1)
		fmt.Println("具体信息:", p)
	}
}

func getAllRoom(msgContent []byte, conn net.Conn) {
	msgUnmarshalled := &roomMessage.C2GS_RoomGet{}
	proto.Unmarshal(msgContent, msgUnmarshalled)

	//要发送给客户端的proto类
	msg2Marshal := &roomMessage.GS2C_RoomGet{}
	//将房间列表写入proto类中

	roomsToDisplayNum := 0
	for i := 0; i < len(rooms); i++ {
		//房间信息初始化
		if rooms[i].isReady || rooms[i].isGameInProgress {
			continue
		}
		roomsToDisplayNum++
	}

	msg2Marshal.Rooms = make([]*roomMessage.Room, roomsToDisplayNum)
	i := 0
	k := 0
	for ; i < len(rooms); i++ {
		//房间信息初始化
		if rooms[i].isReady || rooms[i].isGameInProgress {
			continue
		}

		msg2Marshal.Rooms[k] = new(roomMessage.Room)
		msg2Marshal.Rooms[k].RoomId = rooms[i].roomId
		msg2Marshal.Rooms[k].Players = make([]*roomMessage.Player, len(rooms[i].players))
		for j := 0; j < len(rooms[i].players); j++ {
			msg2Marshal.Rooms[k].Players[j] = new(roomMessage.Player)
			msg2Marshal.Rooms[k].Players[j].PlayerId = rooms[i].players[j].playerId
			msg2Marshal.Rooms[k].Players[j].PlayName = rooms[i].players[j].playerName
			msg2Marshal.Rooms[k].Players[j].PlayerStatus = int32(rooms[i].players[j].status)
		}
		k++
	}
	//消息序列化
	msgContent2Send, _ := proto.Marshal(msg2Marshal)
	msg2Send := TC_Combine(messageType.S_ROOMGET, msgContent2Send)
	sendMsg2One(conn, msg2Send)
}

//为用户创建房间，并告知其房间信息
func crtRoom(msgContent []byte, conn net.Conn) {
	msgUnmarshalled := &roomMessage.C2GS_RoomCrt{}
	proto.Unmarshal(msgContent, msgUnmarshalled)

	// roomCrtor := Player{
	// 	&conn,
	// 	msgUnmarshalled.GetRoomOwner().GetPlayerId(),
	// 	msgUnmarshalled.GetRoomOwner().GetPlayName(),
	// }
	var roomCrtor *Player

	//从维护的玩家队列中找到该玩家信息
	PlayersArrMutex.RLock()
	for i, p := range playersArr {
		if p.playerId == msgUnmarshalled.GetRoomOwner().GetPlayerId() {
			roomCrtor = &playersArr[i]
		}
	}
	PlayersArrMutex.RUnlock()
	newRoom := crtRoomByPlayer(roomCrtor)

	//加锁
	RoomsMutex.Lock()
	pointer2NewRoom := new(GameRoom)
	pointer2NewRoom = &newRoom
	rooms = append(rooms, pointer2NewRoom)
	//pointer2NewRoom := &rooms[len(rooms)-1]
	RoomsMutex.Unlock()

	//要发送给客户端的proto类
	msg2Marshal := &roomMessage.GS2C_RoomCrt{}
	//
	msg2Marshal.ResStatus = 100
	msg2Marshal.RoomId = newRoom.roomId
	//消息序列化
	msgContent2Send, _ := proto.Marshal(msg2Marshal)
	msg2Send := TC_Combine(messageType.S_ROOMCRT, msgContent2Send)
	sendMsg2One(conn, msg2Send)

	//该客户端进入房间，关闭该线程，进入房间内逻辑线程
	go pointer2NewRoom.roomGameServerStart()
	runtime.Goexit()
}

//用户加入房间处理
func joinRoom(msgContent []byte, conn net.Conn) {
	var isJoined bool = false
	msgUnmarshalled := &roomMessage.C2GS_RoomJoin{}
	proto.Unmarshal(msgContent, msgUnmarshalled)

	pid := msgUnmarshalled.GetPlayerId()
	rid := msgUnmarshalled.GetRoomId()

	msg2Marshal := &roomMessage.GS2C_RoomJoin{}

	//给PlayersArrMutex加读锁
	PlayersArrMutex.RLock()
	RoomsMutex.RLock()
	defer PlayersArrMutex.RUnlock()
	defer RoomsMutex.RUnlock()
	for subscript, groom := range rooms {
		if groom.roomId == rid {
			//若房间人数已满，则加入失败
			if len(groom.players) == STANDARD_PLAYER_IN_ROOM {
				msg2Marshal.ResStatus = 301
				break
			}
			for index, playerJoin := range playersArr {
				if pid == playerJoin.playerId {
					//改变客户端信息
					playersArr[index].status = playerstatus.NotReady
					playersArr[index].roomId = groom.roomId

					//将指向客户端队列中该用户的指针放进房间内
					if len(rooms[subscript].players) < STANDARD_PLAYER_IN_ROOM {
						rooms[subscript].playersMutex.Lock()
						if len(rooms[subscript].players) < STANDARD_PLAYER_IN_ROOM {
							rooms[subscript].players = append(groom.players, &playersArr[index])
							isJoined = true
						}
						rooms[subscript].playersMutex.Unlock()
					}

					//打印
					for _, r := range rooms {
						for _, p := range r.players {
							fmt.Println(r.roomId, " ", *p)
						}
					}
					//编辑回复消息
					msg2Marshal.ResStatus = 100
					msg2Marshal.RoomInfo = &roomMessage.Room{}
					msg2Marshal.RoomInfo.RoomId = rid
					//将房间内客户端信息写入消息中,锁
					rooms[subscript].playersMutex.RLock()
					for _, playersInRoom := range rooms[subscript].players {
						tplayer := &roomMessage.Player{
							PlayerId:     playersInRoom.playerId,
							PlayName:     playersInRoom.playerName,
							PlayerStatus: int32(playersInRoom.status),
						}
						msg2Marshal.RoomInfo.Players = append(msg2Marshal.RoomInfo.Players, tplayer)
					}
					rooms[subscript].playersMutex.RUnlock()
					// for q := 0; q < len(msg2Marshal.RoomInfo.Players); q++ {
					// 	fmt.Println(msg2Marshal.RoomInfo.GetPlayers()[q].GetPlayerId())
					// 	fmt.Println(msg2Marshal.RoomInfo.GetPlayers()[q].GetPlayerStatus())
					// }
				}
			}
		}
	}
	fmt.Println(rooms)

	//若该客户端成功加入房间,则关闭当前线程，
	if isJoined {
		msgContent2Send, _ := proto.Marshal(msg2Marshal)
		msg2Send := TC_Combine(8, msgContent2Send)
		sendMsg2AllInRoom(msg2Send, rid)
		//向管道中写入信息，通知房间处理线程有客户端加入
		ComInterGorout <- notice.Notice{
			NoticeType: notice.ClientJoin,
			RoomId:     rid,
			PlayerId:   pid,
		}
		//关闭该线程
		runtime.Goexit()
	} else { //若加入失败，向请求客户端发送信息
		msgContent2Send, _ := proto.Marshal(msg2Marshal)
		msg2Send := TC_Combine(messageType.S_ROOMJOIN, msgContent2Send)
		sendMsg2One(conn, msg2Send)
	}

}

func quickJoin(msgContent []byte, conn net.Conn) {
	quickJoinMsgRcv := &roomMessage.C2GS_QuickJoin{}
	proto.Unmarshal(msgContent, quickJoinMsgRcv)
	pid := quickJoinMsgRcv.GetPlayerId()

	isJoined, isCrted := false, false

	var player2QuickJoin *Player
	//在playersArr 中找到这一玩家,锁
	PlayersArrMutex.RLock()
	for pindex, player := range playersArr {
		if pid == player.playerId {
			player2QuickJoin = &playersArr[pindex]
		}
	}
	PlayersArrMutex.RUnlock()

	RoomsMutex.Lock()
	defer RoomsMutex.Unlock()
	//首先尝试加入一个队伍，遍历房间找出房间人最多的但又不超过的第一个房间
	i, j := 1, 0
	for ; i < len(rooms); i++ {
		if len(rooms[i].players) > len(rooms[j].players) {
			j = i
		}
		if len(rooms[j].players) == STANDARD_PLAYER_IN_ROOM-1 {
			break
		}
	}

	//如果房间数量不为0，则一定会加入房间
	if len(rooms) > 0 {
		if len(rooms[j].players) < STANDARD_PLAYER_IN_ROOM {
			rooms[j].playersMutex.Lock()
			if len(rooms[j].players) < STANDARD_PLAYER_IN_ROOM {
				rooms[j].players = append(rooms[j].players, player2QuickJoin)
				isJoined = true
			}
			rooms[j].playersMutex.Unlock()
		}
	}

	//若未加入，则创建
	pointer2NewRoom := new(GameRoom)
	if !isJoined {
		newRoom := crtRoomByPlayer(player2QuickJoin)
		//加入rooms，加锁
		RoomsMutex.Lock()
		pointer2NewRoom = &newRoom
		rooms = append(rooms, pointer2NewRoom)
		RoomsMutex.Unlock()
		isCrted = true
	}

	var msgBytes []byte
	//发送消息
	if isJoined {
		//读锁
		rooms[j].playersMutex.RLock()
		joinedMsg2Send := &roomMessage.GS2C_RoomJoin{}
		joinedMsg2Send.ResStatus = 100
		joinedMsg2Send.RoomInfo = &roomMessage.Room{}
		joinedMsg2Send.RoomInfo.RoomId = rooms[j].roomId
		//将房间内客户端信息写入消息中
		for _, playersInRoom := range rooms[j].players {
			tplayer := &roomMessage.Player{
				PlayerId:     playersInRoom.playerId,
				PlayName:     playersInRoom.playerName,
				PlayerStatus: int32(playersInRoom.status),
			}
			joinedMsg2Send.RoomInfo.Players = append(joinedMsg2Send.RoomInfo.Players, tplayer)
		}
		//序列化
		msgBytes, _ = proto.Marshal(joinedMsg2Send)
		//加消息类型
		msgBytes = TC_Combine(messageType.S_ROOMJOIN, msgBytes)
		//发
		sendMsg2AllInRoom(msgBytes, rooms[j].roomId)
		//取锁
		rooms[j].playersMutex.RUnlock()

	} else if isCrted {
		crtedMsg2Send := &roomMessage.GS2C_RoomCrt{
			ResStatus: 100,
			RoomId:    pointer2NewRoom.roomId,
		}
		msgBytes, _ = proto.Marshal(crtedMsg2Send)
		msgBytes = TC_Combine(messageType.S_ROOMCRT, msgBytes)
		//发
		sendMsg2One(conn, msgBytes)
	}
}

//根据用户信息创建一个新房间
func crtRoomByPlayer(roomOwner *Player) (gr GameRoom) {
	rand.Seed(time.Now().UnixNano())
	gr = GameRoom{
		roomId:           generateUniqueId(),
		isGameInProgress: false,
		isReady:          false,
		allMsg:           map[*Player][][]byte{},
		currentFrame:     0,
		playerRole:       map[*Player]byte{},
		msgMutex:         new(sync.RWMutex),
		playersMutex:     new(sync.RWMutex),
		roleMutex:        new(sync.RWMutex),
	}
	//gr.messages = make(chan *operation.C2GS_Operation, 32)
	gr.msgCom = make(chan InnerGameNotice, 1)
	//改变客户端信息
	roomOwner.status = playerstatus.NotReady
	roomOwner.roomId = gr.roomId
	gr.players = append(gr.players, roomOwner)
	return gr
}

func appCloseOuterRoom(msgContent []byte) {
	clientCloseMsgRcv := &appclose.C2GS_AppClose{}
	proto.Unmarshal(msgContent, clientCloseMsgRcv)

	//将其从playersArr中移出，加锁
	PlayersArrMutex.Lock()
	defer PlayersArrMutex.Unlock()
	for j := 0; j < len(playersArr); j++ {
		if playersArr[j].playerId == clientCloseMsgRcv.GetPlayerId() {
			//playersArr[j].conn.Close()
			playersArr = append(playersArr[:j], playersArr[j+1:]...)
			break
		}
	}

	fmt.Printf("目前还剩%d位玩家在线\n", len(playersArr))
	for i, p := range playersArr {
		fmt.Printf("第%d位用户\n:", i)
		fmt.Println("具体信息:", p)
	}

	runtime.Goexit()
}

func generateUniqueId() int32 {
	idLock.Lock()
	UniqueId++
	idLock.Unlock()
	return UniqueId
}
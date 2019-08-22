package roomManage

import (
	"fmt"
	"goproto/roomMessage"
	"io"
	"math/rand"
	"net"
	"notice"
	"os"
	"playerstatus"
	"runtime"
	"serverlog"
	"time"

	"github.com/golang/protobuf/proto"
)

//日志记录器
var slog serverlog.ServerLog

//所有玩家列表
var playersArr []Player

//所有房间列表
var rooms []GameRoom

//用于在process中通知房间逻辑线程有客户端加入等
//noticeType:1-客户端加入
//noticeType:2-游戏开始
//noticeType:3-房间解散
var comInterGorout chan notice.Notice

//status 0：未在房间中
//1：未就绪
//2：已准备
//3：选好英雄
//4：游戏进行中
type Player struct {
	conn       net.Conn
	playerId   int32
	playerName string
	status     byte
	roomId     int32
}

type GameRoom struct {
	roomId           int32
	players          []*Player
	isGameInProgress bool
	message          [][]byte
	currentFrame     int32
	playerRole       map[*Player]byte
	readyPlayerNum   byte
}

func (gr *GameRoom) roomGameServerStart() {
	//为房间中每个客户端开创一个线程，对其发送消息进行接受
	for i := 0; i < len(gr.players); i++ {
		go gr.msgReceive(gr.players[i])
	}

	//如果没在游戏中
	for {
		if !gr.isGameInProgress {
			for {
				select {
				case noticeMsg := <-comInterGorout:
					fmt.Println(noticeMsg.PlayerId, noticeMsg.RoomId, noticeMsg.NoticeType)
					//如果有玩家加入某一个房间
					if noticeMsg.NoticeType == 1 {
						// fmt.Println("有玩家加入", noticeMsg.PlayerId, noticeMsg.RoomId, noticeMsg.NoticeType)
						// for _, p := range gr.players {
						// 	fmt.Println(*p)
						// }
						//若有玩家加入本房间，则为其开创新协程
						if noticeMsg.RoomId == gr.roomId {
							for i := 0; i < len(gr.players); i++ {
								if noticeMsg.PlayerId == gr.players[i].playerId {
									fmt.Println("为新玩家开创线程")
									go gr.msgReceive(gr.players[i])
								}
							}
						}
					}
					//游戏开始
					if noticeMsg.NoticeType == 2 {

					}
					//所有玩家退出房间
					if noticeMsg.NoticeType == 3 {
						for i := 0; i < len(rooms); i++ {
							if rooms[i].roomId == gr.roomId {
								rooms = append(rooms[:i], rooms[i+1:]...)
							}
						}
						runtime.Goexit()
					}
				default:
					continue
				}
			}
		} else {

		}
	}

}

func (gr *GameRoom) msgReceive(pplayer *Player) {

	//接收消息
	for {
		buf := make([]byte, 128)
		msgLen, err := pplayer.conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Log2file(err.Error())
			fmt.Println("Receive Failed, err:", err)
		}
		msgType := buf[0]
		msgContent := buf[1:msgLen]

		// fmt.Println("收到房间内消息：", msgType)
		// for _, r := range rooms {
		// 	fmt.Println("所有房间")
		// 	fmt.Println(r)
		// 	for _, p := range r.players {
		// 		fmt.Println(*p)
		// 	}
		// }
		//游戏尚未开始逻辑
		if !gr.isGameInProgress {
			switch msgType {
			//退出房间
			case 9:
				fmt.Println("退出房间")
				gr.quitRoom(msgContent)
			//玩家准备
			case 11:
				fmt.Println("玩家准备")
				gr.playerPrepare(msgContent)
			//进入选择
			case 16:
				fmt.Println("进入选择")
				gr.startSelect(msgContent)
			//英雄选择
			case 13:
				fmt.Println("英雄选择")
				gr.roleSelect(msgContent)
			}
		} else { //游戏进行中消息收发逻辑

		}
	}
}

func (gr *GameRoom) quitRoom(msgContent []byte) {
	var player2Quit *Player
	var quitRes bool = false

	roomQuitMsgRcv := &roomMessage.C2GS_RoomQuit{}
	proto.Unmarshal(msgContent, roomQuitMsgRcv)

	for i := 0; i < len(gr.players); i++ {
		if roomQuitMsgRcv.GetPlayerId() == gr.players[i].playerId {
			player2Quit = gr.players[i]
			//判断其状态
			if player2Quit.status == playerstatus.Gaming {
				break
			}
			//将其移出房间，转入房间外协程
			go process(player2Quit.conn)

			//从房间中将该玩家指针删除, 加锁 lock
			gr.players = append(gr.players[:i], gr.players[i+1:]...)
			//改变其状态
			player2Quit.status = playerstatus.NotInRoom
			player2Quit.roomId = 0
			quitRes = true
			break
		}
	}

	roomQuitMsg2Send := &roomMessage.GS2C_RoomQuit{}
	roomQuitMsg2Send.PlayerId = roomQuitMsgRcv.GetPlayerId()
	roomQuitMsg2Send.ResStatus = 100
	if !quitRes {
		roomQuitMsg2Send.ResStatus = 301
	}

	msg2Send, err := proto.Marshal(roomQuitMsg2Send)
	if err != nil {
		fmt.Println("房间退出信息序列化失败")
		os.Exit(-1)
	}
	msg2Send = TC_Combine(10, msg2Send)
	//如果退出成功向退出客户端发消息，再向所有房间内客户端发送消息
	//如果退出失败只向退出客户端发消息
	sendMsg2One(player2Quit.conn, msg2Send)
	if quitRes {

		//
		if len(gr.players) == 0 {
			comInterGorout <- notice.Notice{
				NoticeType: 3,
				RoomId:     gr.roomId,
			}
		} else {
			sendMsg2AllInRoom(msg2Send, gr.roomId)
		}
		//退出该协程
		runtime.Goexit()
	}
}

func (gr *GameRoom) playerPrepare(msgContent []byte) {
	var isReady bool = false
	playerReadyMsgRcv := &roomMessage.C2GS_PlayerReady{}
	proto.Unmarshal(msgContent, playerReadyMsgRcv)

	for i := 0; i < len(gr.players); i++ {
		if playerReadyMsgRcv.GetPlayerId() == gr.players[i].playerId {
			gr.players[i].status = playerstatus.IsReady
			isReady = true
			break
		}
	}

	//回复客户端
	playerReadyMsg2Send := &roomMessage.GS2C_PlayerReady{}
	playerReadyMsg2Send.PlayerId = playerReadyMsgRcv.GetPlayerId()
	playerReadyMsg2Send.ResStatus = 501
	if isReady {
		playerReadyMsg2Send.ResStatus = 100
	}
	//序列化
	msg2Send, err := proto.Marshal(playerReadyMsg2Send)
	if err != nil {
		fmt.Println("玩家准备消息序列化失败")
	}
	msg2Send = TC_Combine(12, msg2Send)
	for _, pppp := range gr.players {
		fmt.Println(pppp.playerId)
		fmt.Println(pppp.playerName)
		fmt.Println(pppp.conn)
	}
	sendMsg2AllInRoom(msg2Send, gr.roomId)
}

func (gr *GameRoom) startSelect(msgContent []byte) {
	var ableStart bool = false
	var monster *Player
	startSelectMsgRcv := &roomMessage.C2GS_StartChoose{}
	proto.Unmarshal(msgContent, startSelectMsgRcv)

	startSelectMsg2Send := &roomMessage.GS2C_StartChoose{}

	if startSelectMsgRcv.GetRoomId() == gr.roomId {
		if len(gr.players) == 5 {
			for _, player2Start := range gr.players {
				if startSelectMsgRcv.GetPlayerId() == player2Start.playerId {
					//随机生成怪物
					rand.Seed(time.Now().UnixNano())
					monster = gr.players[rand.Intn(len(gr.players))]

					//
					gr.playerRole[monster] = 0

					ableStart = true
				}
			}
		}
	}

	if ableStart {
		startSelectMsg2Send.Ltype = 0
		startSelectMsg2Send.PlayerId = monster.playerId
		startSelectMsg2Send.RoomId = gr.roomId
	}

	//序列化
	msg2Send, _ := proto.Marshal(startSelectMsg2Send)

	msg2Send = TC_Combine(17, msg2Send)
	sendMsg2AllInRoom(msg2Send, gr.roomId)
}

func (gr *GameRoom) roleSelect(msgContent []byte) {
	var player2SelectRole *Player
	var isSelected bool = false
	roleSelectMsgRecv := &roomMessage.C2GS_ChooseLegend{}
	proto.Unmarshal(msgContent, roleSelectMsgRecv)

	if roleSelectMsgRecv.GetRoomId() == gr.roomId {
		for _, player := range gr.players {
			if roleSelectMsgRecv.GetPlayerId() == player.playerId {
				player2SelectRole = player
				gr.playerRole[player] = byte(roleSelectMsgRecv.GetLtype())
				isSelected = true
				break
			}
		}
	}

	roleSelectMsg2Send := &roomMessage.GS2C_ChooseLegend{}
	roleSelectMsg2Send.PlayerId = roleSelectMsgRecv.GetPlayerId()
	roleSelectMsg2Send.Ltype = roleSelectMsgRecv.GetLtype()
	roleSelectMsg2Send.ResStatus = 601
	//如果选择失败
	if !isSelected {
		msg2Send, _ := proto.Marshal(roleSelectMsg2Send)
		msg2Send = TC_Combine(14, msg2Send)
		sendMsg2One(player2SelectRole.conn, msg2Send)
	} else { //成功
		roleSelectMsg2Send.ResStatus = 100
		msg2Send, _ := proto.Marshal(roleSelectMsg2Send)
		msg2Send = TC_Combine(14, msg2Send)
		sendMsg2AllInRoom(msg2Send, gr.roomId)
		//需要对gr.readyPlayerNum加锁，后续考虑
		gr.readyPlayerNum++
		//如果所有玩家准备完成，告知客户端开始游戏
		if gr.readyPlayerNum == 4 {
			gr.isGameInProgress = true
			//消息
			gameStartMsg2Send := &roomMessage.GS2C_GameStart{}
			gameStartMsg2Send.RoomId = gr.roomId
			//序列化
			msg2Send2, _ := proto.Marshal(gameStartMsg2Send)
			msg2Send = TC_Combine(15, msg2Send)
			sendMsg2AllInRoom(msg2Send2, gr.roomId)
		}
	}

}

func process(conn net.Conn) error {
	fmt.Println("Client %v , %v Is Accepted!", conn.RemoteAddr().String(), conn.LocalAddr().String())

	for {
		buf := make([]byte, 128)
		msgLen, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Log2file(err.Error())
			fmt.Println("Receive Failed, err:", err)
			return err
		}
		msgType := buf[0]
		msgContent := buf[1:msgLen]
		fmt.Println("收到房间外消息：", msgType)
		for _, r := range rooms {
			fmt.Println("所有房间")
			fmt.Println(r)
			for _, p := range r.players {
				fmt.Println(*p)
			}
		}

		switch msgType {
		//玩家登录
		case 1:
			playerEnterGame(msgContent, conn)
		//房间列表获取
		case 3:
			getAllRoom(msgContent, conn)
		//房间创建
		case 5:
			crtRoom(msgContent, conn)
		//房间加入
		case 7:
			joinRoom(msgContent, conn)
		}
	}

	return nil
}

func TC_Combine(msgTye byte, sndMsgContent []byte) []byte {
	var buffer bytes.Buffer
	sendMsgType := []byte{msgTye}
	buffer.Write(sendMsgType)
	buffer.Write(sndMsgContent)
	msg2Send := buffer.Bytes()
	return msg2Send
}

func sendMsg2All(msg []byte) {
	for _, player := range playersArr {
		_, e := sendMsg2One(player.conn, msg)
		if e != nil {
			slog.Log2file(e.Error())
			continue
		}
	}
}

func sendMsg2AllInRoom(msg []byte, roomId int32) {
	fmt.Println("正在发送")
	for _, room := range rooms {
		if room.roomId == roomId {
			for _, player := range room.players {
				fmt.Println("id", player.playerId)
				_, e := sendMsg2One(player.conn, msg)
				if e != nil {
					//slog.Log2file(e.Error())
					fmt.Println(e.Error())
					continue
				}
			}
		}
	}
}

func sendMsg2One(clientConn net.Conn, msg []byte) (n int, e error) {
	fmt.Println("正在针对单人发送")
	for _, p := range playersArr {
		fmt.Println(p)
	}
	fmt.Println("房间内")
	for _, r := range rooms {
		for _, pp := range r.players {
			fmt.Println(*pp)
		}
	}
	n, e = clientConn.Write(msg)
	return n, e
}

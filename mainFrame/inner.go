package mainFrame

import (
	"fmt"
	"goproto/appclose"
	"goproto/operation"
	"goproto/roomMessage"
	"io"
	"math/rand"
	"messageType"
	"notice"
	"playerstatus"
	"runtime"
	"time"

	"github.com/golang/protobuf/proto"
)

func (gr *GameRoom) roomGameServerStart() {
	fmt.Printf("房间%d线程启动\n", gr.roomId)

	//为房间中每个客户端开创一个线程，对其发送消息进行接受
	for i := 0; i < len(gr.players); i++ {
		go gr.msgReceive(gr.players[i])
	}

	//如果没在游戏中
	for {
		if !gr.isGameInProgress {
		BEGIN:
			for {
				select {
				case noticeMsg := <-ComInterGorout:
					fmt.Println(noticeMsg.PlayerId, noticeMsg.RoomId, noticeMsg.NoticeType)

					//如果有玩家加入某一个房间
					if noticeMsg.NoticeType == notice.ClientJoin {
						fmt.Println("有玩家加入", noticeMsg.PlayerId, noticeMsg.RoomId, noticeMsg.NoticeType)
						for _, p := range gr.players {
							fmt.Println(*p)
						}
						//若有玩家加入本房间，则为其开创新协程
						if noticeMsg.RoomId == gr.roomId {
							for i := 0; i < len(gr.players); i++ {
								if noticeMsg.PlayerId == gr.players[i].playerId {
									fmt.Println("为新玩家开创线程")
									go gr.msgReceive(gr.players[i])
									break
								}
							}
							fmt.Println("房间内玩家人数：", len(gr.players))
						}
					}

					if noticeMsg.NoticeType == notice.RoomDismiss {
						if noticeMsg.RoomId == gr.roomId {
							RoomsMutex.Lock()
							defer RoomsMutex.Unlock()
							for i := 0; i < len(rooms); i++ {
								if rooms[i].roomId == gr.roomId {
									rooms = append(rooms[:i], rooms[i+1:]...)
									break
								}
							}
							fmt.Println("房间数：", len(rooms))
							runtime.Goexit()
						}
					}
				case innerGameMsg := <-gr.msgCom:
					//游戏开始
					if innerGameMsg.noticeType == notice.GameStart {
						break BEGIN
					}
					//房间解除，加锁
					if innerGameMsg.noticeType == notice.RoomDismiss {
						RoomsMutex.Lock()
						defer RoomsMutex.Unlock()
						for i := 0; i < len(rooms); i++ {
							if rooms[i].roomId == gr.roomId {
								rooms = append(rooms[:i], rooms[i+1:]...)
								break
							}
						}
						fmt.Println("房间数：", len(rooms))
						runtime.Goexit()
					}
					//游戏结束
					if innerGameMsg.noticeType == notice.GameEnd {
						gameEnded := false
						for _, player := range gr.players {
							if player.status != playerstatus.NotInRoom {
								break
							}
							gameEnded = true
						}

						if gameEnded {
							RoomsMutex.Lock()
							defer RoomsMutex.Unlock()
							for i := 0; i < len(rooms); i++ {
								if rooms[i].roomId == gr.roomId {
									rooms = append(rooms[:i], rooms[i+1:]...)
									break
								}
							}
							runtime.Goexit()
						}
					}

				default:
					continue
				}
			}
		} else { //游戏进行中逻辑
			fmt.Println("开始发游戏消息")
			ticker := time.NewTicker(time.Millisecond * 50)
			for {
				<-ticker.C
				//从消息队列中取出消息，发给客户端,加锁
				if len(gr.emptyMsg) != 0 {
					gr.emptyMutex.Lock()
					suppleFrameMsgSnd := &operation.GS2C_SuppleFrame{}
					msg2Send, gameMarerr := proto.Marshal(suppleFrameMsgSnd)
					if gameMarerr != nil {
						Slog.Log2filef("补帧操作序列化失败：%s", gameMarerr.Error())
						continue
					}
					//清空
					gr.emptyMsg = append([]*operation.C2GS_SuppleFrame{})
					gr.emptyMutex.Unlock()
					//添加消息类型
					msg2Send = TC_Combine(messageType.S_SUPPLEFRAME, msg2Send)
					//发送
					gr.sendMsg2AllInThisRoom(msg2Send)
					continue
				}

				if len(gr.message) != 0 {
					operationMsg := &operation.GS2C_Operation{}
					gr.frameMutex.Lock()
					operationMsg.ServerFrame = gr.currentFrame
					gr.currentFrame++
					gr.frameMutex.Unlock()
					gr.msgMutex.Lock()
					//取出消息
					for _, msg := range gr.message {
						operationMsg.Operations = append(operationMsg.Operations, msg)
					}
					//清空message
					gr.message = append([]*operation.C2GS_Operation{})

					gr.msgMutex.Unlock()
					//序列化
					msg2Send, gameMarerr := proto.Marshal(operationMsg)
					if gameMarerr != nil {
						Slog.Log2filef("游戏操作消息序列化失败：%s", gameMarerr.Error())
						continue
					}
					//添加消息类型
					msg2Send = TC_Combine(messageType.S_OPERATION, msg2Send)
					//发送
					gr.sendMsg2AllInThisRoom(msg2Send)

				}
			}
			ticker.Stop()
		}
	}

}

func (gr *GameRoom) msgReceive(pplayer *Player) {

	//接收消息
	rcvErrCount := 0
	for {
		var msglen int32 = 0

		if pplayer.status == playerstatus.Deleted {
			runtime.Goexit()
		}

		//读取消息长度
		var msgLenByte []byte
		var lenLenHaveRead int32 = 0
		for {
			tempBuf := make([]byte, 4-lenLenHaveRead)
			lenByte, e := pplayer.conn.Read(tempBuf[0 : 4-lenLenHaveRead])

			if e != nil {
				if e == io.EOF {
					pplayer.conn.Close()
					runtime.Goexit()
				}
				Slog.Log2file(e.Error())
				fmt.Println("Receive Failed, err:", e)
				rcvErrCount++
				if rcvErrCount >= 20 {
					pplayer.conn.Close()
					rcvErrCount = 0
					runtime.Goexit()
				}
			}
			msgLenByte = CombineBytes(msgLenByte, tempBuf)
			lenLenHaveRead += int32(lenByte)
			if lenLenHaveRead == 4 {
				break
			}
		}
		msglen = BytesToInt(msgLenByte)
		//fmt.Println("消息长度：", msglen)
		if msglen < 1 || msglen > 128 {
			var flush []byte
			pplayer.conn.Read(flush)
			continue
		}

		var buf []byte
		//已经读到的字节数
		var protoLenHaveRead int32 = 0
		for {
			tempBuf := make([]byte, msglen-protoLenHaveRead)
			protoLen, err := pplayer.conn.Read(tempBuf[0 : msglen-protoLenHaveRead])

			if err != nil {
				if err == io.EOF {
					pplayer.conn.Close()
					runtime.Goexit()
				}

				Slog.Log2file(err.Error())
				fmt.Println("Receive Failed, err:", err)
				rcvErrCount++
			}
			buf = CombineBytes(buf, tempBuf)
			protoLenHaveRead += int32(protoLen)
			if protoLenHaveRead == msglen {
				break
			}
		}

		//在游戏状态中时，若该客户端在一定时间内未接受到任何消息,则断开连接，关闭协程

		if len(buf) < 1 || len(buf) > 128 {
			continue
		}

		//fmt.Println("房间内收到消息")

		msgType := buf[0]
		msgContent := buf[1:msglen]

		if msgType != messageType.C_OPERATION {
			fmt.Println("收到的消息类型是：", msgType)
		}

		if msgType == messageType.C_APPCLOSE {
			fmt.Println("客户端退出游戏")
			gr.clientAppClose(msgContent, pplayer)
		}

		//////////////////////////////////////////////

		//游戏尚未开始逻辑
		if !gr.isGameInProgress {
			switch msgType {
			//退出房间
			case messageType.C_ROOMQUIT:
				fmt.Println("退出房间")
				gr.quitRoom(msgContent, pplayer)
			//玩家准备
			case messageType.C_PLAYERREADY:
				fmt.Println("玩家准备")
				gr.playerPrepare(msgContent, pplayer)
			//进入选择
			case messageType.C_STARTCHOOSE:
				fmt.Println("进入选择")
				gr.startSelect(msgContent, pplayer)
			//英雄选择
			case messageType.C_CHOOSELEGEND:
				fmt.Println("英雄选择")
				gr.roleSelect(msgContent, pplayer)
			// case messageType.C_APPCLOSE:
			// 	fmt.Println("客户端退出游戏")
			// 	gr.clientAppClose(msgContent)
			default:
				var flush []byte
				pplayer.conn.Read(flush)
			}

		} else { //游戏进行中消息收发逻辑

			switch msgType {
			//客户端同步信息
			case messageType.C_PLAYERSTATUSSYNC:
				fmt.Println("客户端同步")
				gr.playerStatusSync(msgContent, pplayer)
				continue
			//补帧请求
			case messageType.C_SUPPLEFRAME:
				suppleFrameMsg := &operation.C2GS_SuppleFrame{}
				umerr := proto.Unmarshal(msgContent, suppleFrameMsg)
				if umerr != nil {
					Slog.Log2filef("房间%d游戏内消息解析失败\n", gr.roomId)
					Slog.Log2filef("错误详细信息：%s", umerr.Error())
					var flush []byte
					pplayer.conn.Read(flush)
					continue
				}
				gr.emptyMutex.Lock()
				gr.emptyMsg = append(gr.emptyMsg, suppleFrameMsg)
				gr.emptyMutex.Unlock()
			//游戏操作
			case messageType.C_OPERATION:
				//解析proto
				operationMsgRcv := &operation.C2GS_Operation{}
				umerr := proto.Unmarshal(msgContent, operationMsgRcv)

				if umerr != nil {
					Slog.Log2filef("房间%d游戏内消息解析失败\n", gr.roomId)
					Slog.Log2filef("错误详细信息：%s", umerr.Error())
					var flush []byte
					pplayer.conn.Read(flush)
					continue
				}

				//将解析出来的消息放进message中
				gr.msgMutex.Lock()
				gr.message = append(gr.message, operationMsgRcv)
				gr.msgMutex.Unlock()

			case messageType.C_GAMEEND:
				pplayer.roomId = 0
				pplayer.status = playerstatus.NotInRoom
				//通知
				gr.msgCom <- InnerGameNotice{
					noticeType: notice.GameEnd,
					playerId:   pplayer.playerId,
				}
				go Process(pplayer.conn)
				runtime.Goexit()

			}
		}
	}
}

func (gr *GameRoom) quitRoom(msgContent []byte, pplayer *Player) {
	gr.playersMutex.Lock()
	defer gr.playersMutex.Unlock()
	var player2Quit *Player
	var quitRes bool = false

	roomQuitMsgRcv := &roomMessage.C2GS_RoomQuit{}
	ume := proto.Unmarshal(msgContent, roomQuitMsgRcv)
	if ume != nil {
		Slog.Log2filef("Quit Room Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	//fmt.Println(roomQuitMsgRcv.GetRoomId())
	//fmt.Println(roomQuitMsgRcv.GetPlayerId())

	if roomQuitMsgRcv.GetRoomId() == gr.roomId {
		//如果房主已点击开始选人，则无法退出
		if !gr.isReady {
			for i := 0; i < len(gr.players); i++ {
				if roomQuitMsgRcv.GetPlayerId() == gr.players[i].playerId {
					player2Quit = gr.players[i]

					//将其移出房间，转入房间外协程
					go Process(player2Quit.conn)

					//从房间中将该玩家指针删除, 加锁 lock
					gr.players = append(gr.players[:i], gr.players[i+1:]...)
					//改变其状态
					player2Quit.status = playerstatus.NotInRoom
					player2Quit.roomId = 0
					//清除其英雄选择数据
					gr.roleMutex.Lock()
					delete(gr.playerRole, player2Quit)
					gr.roleMutex.Unlock()

					//如果玩家已经准备，则房间中已准备数量减一
					if player2Quit.status == playerstatus.IsReady {
						player2Quit.status = playerstatus.NotInRoom
					}
					quitRes = true
					break
				}
			}
		}
	}

	//要发送的proto
	roomQuitMsg2Send := &roomMessage.GS2C_RoomQuit{}
	roomQuitMsg2Send.PlayerId = roomQuitMsgRcv.GetPlayerId()
	roomQuitMsg2Send.ResStatus = 301
	if quitRes {
		roomQuitMsg2Send.ResStatus = 100
	}

	msg2Send, err := proto.Marshal(roomQuitMsg2Send)
	if err != nil {
		fmt.Println("房间退出信息序列化失败")
	}
	msg2Send = TC_Combine(messageType.S_ROOMQUIT, msg2Send)
	//如果退出成功向退出客户端发消息，再向所有房间内客户端发送消息
	//如果退出失败只向退出客户端发消息
	if player2Quit != nil {
		sendMsg2One(player2Quit.conn, msg2Send)
	}

	if quitRes {

		//
		if len(gr.players) == 0 {
			// ComInterGorout <- notice.Notice{
			// 	NoticeType: notice.RoomDismiss,
			// 	RoomId:     gr.roomId,
			// }
			gr.msgCom <- InnerGameNotice{
				noticeType: notice.RoomDismiss,
			}
		} else {
			gr.sendMsg2AllInThisRoom(msg2Send)
			//sendMsg2AllInRoom(msg2Send, gr.roomId)
		}
		//退出该协程
		fmt.Println("协程退出")
		runtime.Goexit()
	}
	//return ume
}

func (gr *GameRoom) playerPrepare(msgContent []byte, pplayer *Player) {
	var isReady bool = false
	playerReadyMsgRcv := &roomMessage.C2GS_PlayerReady{}
	ume := proto.Unmarshal(msgContent, playerReadyMsgRcv)
	if ume != nil {
		Slog.Log2filef("Prepare Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	//锁
	gr.playersMutex.RLock()
	defer gr.playersMutex.RUnlock()
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
	msg2Send = TC_Combine(messageType.S_PLAYERREADY, msg2Send)
	gr.sendMsg2AllInThisRoom(msg2Send)
	//sendMsg2AllInRoom(msg2Send, gr.roomId)
}

func (gr *GameRoom) startSelect(msgContent []byte, pplayer *Player) {
	var ableStart bool = false
	var monster *Player
	startSelectMsgRcv := &roomMessage.C2GS_StartChoose{}
	ume := proto.Unmarshal(msgContent, startSelectMsgRcv)
	if ume != nil {
		Slog.Log2filef("Start Select Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	startSelectMsg2Send := &roomMessage.GS2C_StartChoose{}
	//锁
	gr.roleMutex.Lock()
	gr.playersMutex.RLock()
	if startSelectMsgRcv.GetRoomId() == gr.roomId {
		if gr.getReadyNum() == STANDARD_PLAYER_IN_ROOM {
			//改变房间状态，无法退出
			gr.isReady = true

			for _, player2Start := range gr.players {
				if startSelectMsgRcv.GetPlayerId() == player2Start.playerId {
					//随机生成怪物
					rand.Seed(time.Now().UnixNano())
					monster = gr.players[rand.Intn(len(gr.players))]
					monster.status = playerstatus.Locked

					gr.playerRole[monster] = 0

					ableStart = true
				}
			}
		}
	}
	gr.roleMutex.Unlock()
	gr.playersMutex.RUnlock()

	startSelectMsg2Send.ResStatus = 701
	if ableStart {
		startSelectMsg2Send.Ltype = 0
		startSelectMsg2Send.PlayerId = monster.playerId
		startSelectMsg2Send.RoomId = gr.roomId
		startSelectMsg2Send.ResStatus = 100
	}

	//序列化
	msg2Send, _ := proto.Marshal(startSelectMsg2Send)

	msg2Send = TC_Combine(messageType.S_STARTCHOOSE, msg2Send)
	gr.sendMsg2AllInThisRoom(msg2Send)
	//sendMsg2AllInRoom(msg2Send, gr.roomId)
}

func (gr *GameRoom) roleSelect(msgContent []byte, pplayer *Player) {
	var player2SelectRole *Player
	var isSelected bool = false
	roleSelectMsgRecv := &roomMessage.C2GS_ChooseLegend{}
	ume := proto.Unmarshal(msgContent, roleSelectMsgRecv)
	if ume != nil {
		Slog.Log2filef("Select Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	gr.roleMutex.Lock()
	if roleSelectMsgRecv.GetRoomId() == gr.roomId {
		for _, player := range gr.players {
			if roleSelectMsgRecv.GetPlayerId() == player.playerId {
				player2SelectRole = player
				gr.playerRole[player] = byte(roleSelectMsgRecv.GetLtype())
				player.status = playerstatus.Locked
				isSelected = true
				break
			}
		}
	}
	gr.roleMutex.Unlock()

	roleSelectMsg2Send := &roomMessage.GS2C_ChooseLegend{}
	roleSelectMsg2Send.PlayerId = roleSelectMsgRecv.GetPlayerId()
	roleSelectMsg2Send.Ltype = roleSelectMsgRecv.GetLtype()
	roleSelectMsg2Send.ResStatus = 601
	//如果选择失败
	if !isSelected {
		msg2Send, _ := proto.Marshal(roleSelectMsg2Send)
		msg2Send = TC_Combine(messageType.S_CHOOSELEGEND, msg2Send)
		sendMsg2One(player2SelectRole.conn, msg2Send)
	} else { //成功
		fmt.Println("选择成功")
		roleSelectMsg2Send.ResStatus = 100
		msg2Send, _ := proto.Marshal(roleSelectMsg2Send)
		msg2Send = TC_Combine(messageType.S_CHOOSELEGEND, msg2Send)
		gr.sendMsg2AllInThisRoom(msg2Send)
		//time.Sleep(time.Microsecond * 100)
		//如果所有玩家准备完成，告知客户端开始游戏，通知房间主线程开始游戏
		if gr.getLockedNum() == STANDARD_PLAYER_IN_ROOM {
			gr.isGameInProgress = true
			//消息
			gameStartMsg2Send := &roomMessage.GS2C_GameStart{}
			gameStartMsg2Send.RoomId = gr.roomId
			//序列化
			msg2Send2, _ := proto.Marshal(gameStartMsg2Send)
			fmt.Println("Game start:", messageType.S_GAMESTART)
			msg2Send3 := TC_Combine(messageType.S_GAMESTART, msg2Send2)
			gr.sendMsg2AllInThisRoom(msg2Send3)
			//sendMsg2AllInRoom(msg2Send, gr.roomId)
			//通知，加锁
			gr.msgCom <- InnerGameNotice{
				noticeType: notice.GameStart,
			}
		}
	}

}

/**
** 当有客户端强制退出游戏时
** 当房间中仅剩其一位玩家，将其信息清空，关闭连接，关闭该用户收发消息线程。利用消息通知room server，关闭roomserver线程，并移出该房间
** 当房间还有其它玩家时，仅仅对其操作
**/
func (gr *GameRoom) clientAppClose(msgContent []byte, pplayer *Player) {
	var isRoomDismiss bool = false
	clientCloseMsgRcv := &appclose.C2GS_AppClose{}
	ume := proto.Unmarshal(msgContent, clientCloseMsgRcv)
	if ume != nil {
		Slog.Log2filef("Client Close Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	//锁
	var connDelPlayer Player
	gr.playersMutex.RLock()
	defer gr.playersMutex.RUnlock()
	for i := 0; i < len(gr.players); i++ {
		if gr.players[i].playerId == clientCloseMsgRcv.GetPlayerId() {
			if len(gr.players) == 1 {
				//通知房间主线程，关闭该房间线程
				gr.msgCom <- InnerGameNotice{
					noticeType: notice.RoomDismiss,
				}
				isRoomDismiss = true
			}
			tplayer := gr.players[i]
			connDelPlayer = *tplayer

			//改变房间选人数据
			//锁
			gr.roleMutex.Lock()
			delete(gr.playerRole, tplayer)
			gr.roleMutex.Unlock()
			//从房间中将其移除，加锁
			gr.players = append(gr.players[:i], gr.players[i+1:]...)

			//将其从playersArr中移出，加锁
			PlayersArrMutex.Lock()
			defer PlayersArrMutex.Unlock()
			for j := 0; j < len(playersArr); j++ {
				if playersArr[j].playerId == tplayer.playerId {
					playersArr = append(playersArr[:i], playersArr[i+1:]...)
					break
				}
			}
			break
		}
	}

	//若房间未解散，向房间内所有用户发送有玩家退出消息
	if !isRoomDismiss {
		playerQuitMsg2Snd := &roomMessage.GS2C_RoomQuit{
			ResStatus: 100,
			PlayerId:  clientCloseMsgRcv.GetPlayerId(),
		}

		msg2Send, err := proto.Marshal(playerQuitMsg2Snd)
		if err != nil {
			Slog.Log2file(err.Error())
		}
		msg2Send = TC_Combine(messageType.S_ROOMQUIT, msg2Send)
		gr.sendMsg2AllInThisRoom(msg2Send)
	}

	//关闭其连接
	connDelPlayer.conn.Close()
	runtime.Goexit()
}

func (gr *GameRoom) playerStatusSync(msgContent []byte, pplayer *Player) {
	gr.playersMutex.RLock()
	defer gr.playersMutex.RUnlock()
	statusSyncMsgRcv := &roomMessage.C2GS_PlayerStatusSync{}
	ume := proto.Unmarshal(msgContent, statusSyncMsgRcv)
	if ume != nil {
		Slog.Log2filef("Sync Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	var loadCompleted bool = false

	if statusSyncMsgRcv.GetRoomId() == gr.roomId {
		for _, player := range gr.players {
			if statusSyncMsgRcv.GetPlayerId() == player.playerId {
				player.status = playerstatus.Loaded
				break
			}
		}
		//查看全体玩家是否加载完成，加载完成开始游戏
		for _, player := range gr.players {
			if player.status != playerstatus.Loaded {
				break
			}
			loadCompleted = true
		}
		//若加载成功，则向客户端广播可以开始游戏
		if loadCompleted {
			statusSyncMsg2Send := &roomMessage.GS2C_PlayerStatusSync{
				ResStatus: 100,
				Timestamp: int64(time.Now().Unix()),
			}

			//计时任务线程开启
			go gr.timerSeven()

			msg2Send, _ := proto.Marshal(statusSyncMsg2Send)
			msg2Send = TC_Combine(messageType.S_PLAYERSTATUSSYNC, msg2Send)
			gr.sendMsg2AllInThisRoom(msg2Send)
			//sendMsg2AllInRoom(msg2Send, gr.roomId)
		}
	}
}

func (gr *GameRoom) timerSeven() {
	//游戏开始计时
	gr.gameTimer = time.NewTicker(time.Minute * 7)
	for {
		<-gr.gameTimer.C
		for {
			gr.frameMutex.Lock()
			sevenNoteMsg := &operation.GS2C_Operation{
				ServerFrame: gr.currentFrame,
				SevenNote:   true,
			}
			gr.currentFrame++
			gr.frameMutex.Unlock()
			msgBytes, err := proto.Marshal(sevenNoteMsg)
			if err != nil {
				Slog.Log2filef("TimerSeven Proto Marshal Error : %s\n", err.Error())
				continue
			}
			msgBytes = TC_Combine(messageType.S_OPERATION, msgBytes)
			gr.sendMsg2AllInThisRoom(msgBytes)
			break
		}
		break
	}
	gr.gameTimer.Stop()
}

func (gr *GameRoom) playerQuitGame(msgContent []byte, pplayer *Player) {
	playerQuitGameMsgRcv := &operation.C2GS_PlayerQuit{}
	ume := proto.Unmarshal(msgContent, playerQuitGameMsgRcv)
	if ume != nil {
		Slog.Log2filef("Player Quit Unmarshal Error:%s\n", ume.Error())
		var flush []byte
		pplayer.conn.Read(flush)
	}

	gr.playersMutex.Lock()
	defer gr.playersMutex.Unlock()
	if playerQuitGameMsgRcv.GetRoomId() == gr.roomId {
		for pindex, player := range gr.players {
			if player.playerId == playerQuitGameMsgRcv.GetPlayerId() {
				//修改玩家状态
				player.roomId = 0
				player.status = playerstatus.NotInRoom
				//删除
				gr.players = append(gr.players[:pindex], gr.players[pindex+1:]...)
				//开启房间外线程
				go Process(player.conn)
				runtime.Goexit()
			}
		}
	}
}

/**
将房间中每一个用户的房间ID以及状态改为未在房间中的状态
且需将收发消息的
*/
func (gr *GameRoom) gameCompleted(msgContent []byte, pplayer *Player) {
	// gameEndMsgRcv := &operation.C2GS_GameEnd{}
	// proto.Unmarshal(gameEndMsgRcv, msgContent)
	// rid := gameEndMsgRcv.RoomId

	// if rid == gr.roomId {
	// 	for i := 0; i < len(gr.players); i++ {
	// 		gr.players[i].roomId = 0
	// 		gr.players[i].status = playerstatus.NotInRoom
	// 	}
	// }

}

func (gr *GameRoom) sendMsg2AllInThisRoom(msg []byte) {
	// if msg[4] != 22 {
	// 	fmt.Printf("发送的消息类型是:%d\n", msg[4])
	// }
	//count := 0
	for _, player := range gr.players {
		go sendMsg2One(player.conn, msg)
		// _, e := sendMsg2One(player.conn, msg)
		// count++
		// if e != nil {
		// 	Slog.Log2file(e.Error())
		// 	fmt.Println(e.Error())
		// }
	}
	//fmt.Printf("发送人数:%d\n", count)
}

func (gr GameRoom) getReadyNum() (readyPlayersNum byte) {
	gr.playersMutex.RLock()
	defer gr.playersMutex.RUnlock()
	readyPlayersNum = 0
	for _, player := range gr.players {
		if player.status == playerstatus.IsReady {
			readyPlayersNum++
		}
	}
	return readyPlayersNum
}

func (gr GameRoom) getLockedNum() (lockedPlayersNum byte) {
	gr.playersMutex.RLock()
	defer gr.playersMutex.RUnlock()
	lockedPlayersNum = 0
	for _, player := range gr.players {
		if player.status == playerstatus.Locked {
			lockedPlayersNum++
		}
	}
	return lockedPlayersNum
}

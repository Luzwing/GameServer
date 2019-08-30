package mainFrame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"goproto/operation"
	"io"
	"net"
	"notice"
	"playerstatus"
	"runtime"
	"serverlog"
	"sync"
	"time"
)

var joinTimeToSleep int = 80

var conns []net.Conn

//所有玩家列表
// var playersArr []Player
var playersArr []*Player
var PlayersArrMutex sync.RWMutex

//所有房间列表
var rooms []*GameRoom
var RoomsMutex sync.RWMutex

//日志记录器
var Slog serverlog.ServerLog

//用于在process中通知房间逻辑线程有客户端加入等
//noticeType:1-客户端加入
//noticeType:2-游戏开始
//noticeType:3-房间解散
var ComInterGorout []notice.Notice
var ComInterMutex sync.RWMutex

var RoomsAbleQuickJoin map[int]map[int32]byte
var RoomsAbleQJMutex sync.RWMutex

const STANDARD_PLAYER_IN_ROOM = 2

//用户生成房间id
var UniqueId int32
var idLock sync.Mutex

//用户ID生成
var PlayerId int32
var playerIdLock sync.Mutex

//
var QuickJoinTime int
var QJTimeMutex sync.RWMutex

//代表

//锁

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

type InnerGameNotice struct {
	noticeType byte
	playerId   int32
}

type GameRoom struct {
	roomId           int32
	players          []*Player
	isGameInProgress bool
	isReady          bool //用于在玩家退出时判断其是否可以退出
	message          []*operation.C2GS_Operation
	emptyMsg         []*operation.C2GS_SuppleFrame
	msgCom           chan InnerGameNotice
	allMsg           map[*Player][][]byte
	currentFrame     int32
	supplingFrame    bool
	playerRole       map[*Player]byte
	gameTimer        *time.Ticker
	msgMutex         *sync.RWMutex
	playersMutex     *sync.RWMutex
	roleMutex        *sync.RWMutex
	emptyMutex       *sync.RWMutex
	frameMutex       *sync.RWMutex
}

func Close() {
	for _, p := range playersArr {
		p.conn.Close()
	}
	playersArr = nil
}

func empty(arr []interface{}) bool {
	if len(arr) == 0 {
		return true
	}
	return false
}

func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32 = 0
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return x
}

func Int32ToBytes(i int32) []byte {
	var buf []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

func GetAllRooms() {
	RoomsMutex.RLock()
	defer RoomsMutex.RUnlock()
	for rindex, room := range rooms {
		fmt.Printf("房间%d:\n", rindex+1)
		fmt.Printf("房间ID：%d\n", room.roomId)
		fmt.Printf("房间玩家：\n")
		room.playersMutex.RLock()
		for pindex, player := range room.players {
			fmt.Printf("玩家%d:\n", pindex+1)
			fmt.Printf("玩家ID：%d\n", player.playerId)
			fmt.Printf("玩家姓名:%s\n", player.playerName)
		}
		room.playersMutex.RUnlock()
	}
}

func GetAllPlayers() {
	PlayersArrMutex.RLock()
	defer PlayersArrMutex.RUnlock()

	for pindex, player := range playersArr {
		fmt.Printf("玩家%d:\n", pindex+1)
		fmt.Printf("玩家ID：%d\n", player.playerId)
		fmt.Printf("玩家姓名:%s\n", player.playerName)
		fmt.Printf("玩家房间号：%d\n", player.roomId)
		fmt.Printf("玩家状态：%d\n", player.status)
	}
}

func readMsg(conn net.Conn, buf *[]byte, rcvErrCount *int) (int32, error) {
	//读取消息长度
	var msglen int32 = 0
	var msgLenByte []byte
	var lenLenHaveRead int32 = 0
	for {
		tempBuf := make([]byte, 4-lenLenHaveRead)
		lenByte, e := conn.Read(tempBuf[0 : 4-lenLenHaveRead])

		if e != nil {
			if e == io.EOF {
				conn.Close()
				runtime.Goexit()
			}
			Slog.Log2file(e.Error())
			fmt.Println("Receive Failed, err:", e)
			*rcvErrCount++
			if *rcvErrCount >= 20 {
				conn.Close()
				*rcvErrCount = 0
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
		conn.Read(flush)
		return -1, errors.New("Read Error,message length error")
	}

	//已经读到的字节数
	var protoLenHaveRead int32 = 0
	for {
		tempBuf := make([]byte, msglen-protoLenHaveRead)
		protoLen, err := conn.Read(tempBuf[0 : msglen-protoLenHaveRead])

		if err != nil {
			if err == io.EOF {
				conn.Close()
				runtime.Goexit()
			}

			Slog.Log2file(err.Error())
			fmt.Println("Receive Failed, err:", err)
			*rcvErrCount++
		}
		*buf = CombineBytes(*buf, tempBuf)
		protoLenHaveRead += int32(protoLen)
		if protoLenHaveRead == msglen {
			break
		}
	}
	//在游戏状态中时，若该客户端在一定时间内未接受到任何消息,则断开连接，关闭协程
	if len(*buf) < 1 || len(*buf) > 128 {
		return -1, errors.New("Read Error,buffer exception")
	}

	return msglen, nil
}

func DeleteRoomById(rid int32) {
	RoomsMutex.Lock()
	PlayersArrMutex.Lock()
	defer RoomsMutex.Unlock()
	defer PlayersArrMutex.Unlock()
	for _, room := range rooms {
		if rid == room.roomId {
			room.playersMutex.Lock()
			for _, player := range room.players {
				player.roomId = 0
				player.status = playerstatus.Deleted
				player.conn.Close()
				for pi, p := range playersArr {
					if p.playerId == player.playerId {
						playersArr = append(playersArr[:pi], playersArr[pi+1:]...)
					}
				}
			}
			room.playersMutex.Unlock()
			ComInterMutex.Lock()
			ComInterGorout = append(ComInterGorout, notice.Notice{
				NoticeType: notice.RoomDismiss,
				RoomId:     rid,
				PlayerId:   0,
			})
			ComInterMutex.Unlock()
		}
	}

}

func DeletePlayersById(pid int32) {
	PlayersArrMutex.Lock()
	defer PlayersArrMutex.Unlock()
	for pindex, player := range playersArr {
		if pid == player.playerId {
			player.roomId = 0
			player.status = playerstatus.Deleted

			playersArr = append(playersArr[:pindex], playersArr[pindex+1:]...)

			player.conn.Close()
			fmt.Println("删除成功")
		}
	}

}

// func ClearSlice(s *[]interface{}) {
// 	s = append([]interface{}{})
// }

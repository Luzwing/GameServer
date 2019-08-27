package mainFrame

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goproto/operation"
	"net"
	"notice"
	"playerstatus"
	"serverlog"
	"sync"
	"time"
)

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
var ComInterGorout chan notice.Notice
var ComInterMutex sync.RWMutex

const STANDARD_PLAYER_IN_ROOM = 2

//用户生成房间id
var UniqueId int32
var idLock sync.Mutex

//用户ID生成
var PlayerId int32
var playerIdLock sync.Mutex

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
			ComInterGorout <- notice.Notice{
				NoticeType: notice.RoomDismiss,
				RoomId:     rid,
				PlayerId:   0,
			}
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

package mainFrame

import (
	"bytes"
	"encoding/binary"
	"goproto/operation"
	"net"
	"notice"
	"serverlog"
	"sync"
	"time"
)

var conns []net.Conn

//所有玩家列表
var playersArr []Player
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
	messages         chan *operation.C2GS_Operation
	msgCom           chan InnerGameNotice
	allMsg           map[*Player][][]byte
	currentFrame     int32
	playerRole       map[*Player]byte
	gameTimer        *time.Ticker
	msgMutex         *sync.RWMutex
	playersMutex     *sync.RWMutex
	roleMutex        *sync.RWMutex
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

// func ClearSlice(s *[]interface{}) {
// 	s = append([]interface{}{})
// }

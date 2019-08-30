package notice

import (
	"time"
)

const (
	_ byte = iota
	ClientWillJoin
	ClientJoin
	ClientQuickJoin
	GameStart
	RoomDismiss
	GameEnd
	RoomActive
)

//noticeType:1-客户端加入
//noticeType:2-游戏开始
//noticeType:3-房间解散
type Notice struct {
	NoticeType byte
	RoomId     int32
	PlayerId   int32
	IsAbleJoin *bool
	JoinTime   int
}

type QuickJoinReply struct {
	RoomsAbleJoin []int32
}

func TimeOut(duration time.Duration, ch chan int) {
	for {
		time.Sleep(duration)
		ch <- 1
	}
}

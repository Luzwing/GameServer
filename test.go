// fmt.Println("Message Type: %c", msgType)

// msgPContent := &gopo.C2GS_Operation{}

// lla := proto.Unmarshal(msgContent, msgPContent)
// if lla != nil {
// 	continue
// }

// fmt.Println("Message Content: %d", msgPContent.GetPlayerId())

// //发送
// sendProto := gopo.GS2C_Operation{}
// sendProto.Operations = make([]*gopo.C2GS_Operation, 2)
// sendProto.Operations[0] = msgPContent
// data, _ := proto.Marshal(&sendProto)
// fmt.Println(reflect.TypeOf(data).String())
// fmt.Println(len(data))

// sendData := make([]byte, len(data)+1)
// sendData = append([]byte{2}, data...)

// gygy := &gopo.GS2C_Operation{}
// proto.Unmarshal(sendData[1:len(sendData)], gygy)
// fmt.Println(gygy.GetOperations()[0].String())
// _, ss := conn.Write(sendData)
// if ss == nil {
// 	fmt.Println("Send Successfully")
// }

package main

import (
	"fmt"
	"runtime"
	"time"
)

var hh bool

// var t1 []int
// var t2 []*int

//var j *int

type ss struct {
	b []*sss
}

type sss struct {
	c int
}

var tb []*int

func main() {
	fmt.Println(time.Now().UnixNano() / 1e6)
	b := make([]*int, 10)
	for iiii := 0; iiii < 10; iiii++ {
		hh := iiii
		b[iiii] = &hh
	}
	a := &b[3]
	for iiii := 0; iiii < 10; iiii++ {
		fmt.Println(*b[iiii])
	}
	b = append(b[:1], b[2:]...)
	fmt.Println(b)
	fmt.Println(*a)
}

func a() {
	i := 0
	for i < 10 {
		k := 0
		var j *int
		//j := new(int)
		j = &k
		tb = append(tb, j)
		i++
		k = k + i
	}
}

func b() {
	i := 0
	for i < 10 {
		fmt.Println(*tb[i])
		i++
	}
}

func myunlock(i int) int {
	return i
}

func t1() {
	TimeOut(1e9)
	fmt.Println(1)
	go t2()
	runtime.Goexit()
}

func t2() {
	TimeOut(1e9)
	fmt.Println(2)
	go t1()
	runtime.Goexit()

}

func TimeOut(duration time.Duration) {
	time.Sleep(duration)
	// for {
	// 	time.Sleep(duration)
	// }
}

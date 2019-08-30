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

func tte() *int {
	i := 1
	return &i
}

func huh() {
	for i := 0; i < 10; i++ {
		j := i
		tb = append(tb, &j)
	}
}

func change(t *[]int) {
	*t = append(*t, 1234)
}

func main() {
	var tttt []int
	for i := 0; i < 10; i++ {
		tttt = append(tttt, i)
	}
	change(&tttt)
	fmt.Println(tttt)
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

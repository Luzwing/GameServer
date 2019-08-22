package main

import (
	"fmt"
	"mainFrame"
	"net"
	"notice"
	"os"
	"runtime"
)

func main() {
	//设置CPU多核数
	runtime.GOMAXPROCS(runtime.NumCPU())

	//启动日志记录
	logerr := mainFrame.Slog.LogInit()
	if logerr != nil {
		fmt.Println("LogInit Error:", logerr.Error())
		os.Exit(-1)
	}
	defer mainFrame.Slog.Close()

	//初始化管道
	mainFrame.ComInterGorout = make(chan notice.Notice, 1)

	//
	mainFrame.UniqueId = 1000

	//监听端口号
	listen, err := net.Listen("tcp", ":5050")
	if err != nil {
		mainFrame.Slog.Log2file(err.Error())
		fmt.Println("Listen Failed, err:", err)
		os.Exit(-1)
	}
	defer listen.Close()

	fmt.Println("Server Start!")
	//slog.Log2file("Server Start!")
	for {
		conn, err := listen.Accept()
		if err != nil {
			mainFrame.Slog.Log2file(err.Error())
			fmt.Println("Accpet Failed, err:", err)
			continue
		}
		fmt.Printf("Client %s Is Accepted!\n", conn.RemoteAddr().String())
		go mainFrame.Process(conn)
	}
	defer mainFrame.Close()
}

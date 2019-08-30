package main

import (
	"fmt"
	"mainFrame"
	"net"
	"os"
	"runtime"
)

func no(conn net.Conn) net.Conn {
	return conn
}

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

	// //初始化管道
	// mainFrame.ComInterGorout = make(chan notice.Notice, 1)

	//
	mainFrame.UniqueId = 1000
	mainFrame.PlayerId = 123456
	mainFrame.QuickJoinTime = 0

	//
	mainFrame.RoomsAbleQuickJoin = make(map[int]map[int32]byte)

	//监听端口号
	listen, err := net.Listen("tcp", ":5050")
	if err != nil {
		mainFrame.Slog.Log2file(err.Error())
		fmt.Println("Listen Failed, err:", err)
		os.Exit(-1)
	}
	defer listen.Close()

	go processInput()

	fmt.Println("Server Start!")
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

func processInput() {
	fmt.Println("开始读取输入")
	var line string
	for {
		fmt.Scanln(&line)
		if line == "gp" {
			mainFrame.GetAllPlayers()
		} else if line == "gr" {
			mainFrame.GetAllRooms()
		} else if line == "dp" {
			var pid int32
			_, err := fmt.Scanln(&pid)
			if err == nil {
				mainFrame.DeletePlayersById(pid)
			}
		} else if line == "dr" {
			var rid int32
			_, err := fmt.Scanln(&rid)
			if err == nil {
				mainFrame.DeleteRoomById(rid)
			}
		}
	}
}

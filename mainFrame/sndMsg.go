package mainFrame

import (
	"bytes"
	"fmt"
	"net"
)

func CombineBytes(b1 []byte, b2 []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(b1)
	buffer.Write(b2)
	return buffer.Bytes()
}

func TC_Combine(msgTye byte, sndMsgContent []byte) []byte {
	//获取长度
	var msglen int = len(sndMsgContent) + 1
	//转为[]byte
	msgLenByte := Int32ToBytes(int32(msglen))

	var buffer bytes.Buffer
	sendMsgType := []byte{msgTye}
	buffer.Write(msgLenByte)
	buffer.Write(sendMsgType)
	buffer.Write(sndMsgContent)
	msg2Send := buffer.Bytes()
	return msg2Send
}

func sendMsg2All(msg []byte) {
	for _, player := range playersArr {
		_, e := sendMsg2One(player.conn, msg)
		if e != nil {
			Slog.Log2file(e.Error())
			continue
		}
	}
}

func sendMsg2AllInRoom(msg []byte, roomId int32) {
	for _, room := range rooms {
		if room.roomId == roomId {
			for _, player := range room.players {
				_, e := sendMsg2One(player.conn, msg)
				if e != nil {
					Slog.Log2file(e.Error())
					fmt.Println(e.Error())
				}
			}
		}
	}
}

func sendMsg2One(clientConn net.Conn, msg []byte) (n int, e error) {
	n, e = clientConn.Write(msg)
	return n, e
}

package main

import (
	"DAY03/zinx/znet"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(time.Second)
	//1 直接连接远程服务器，得到conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("net dial err:", err)
		return
	}

	//2 连接调用write写数据
	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.5 client Test Message")))
		if err != nil {
			fmt.Println("package faild")
			break
		}

		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("Write faild")
			break
		}

		msgHead := make([]byte, 8)
		cnt, err := conn.Read(msgHead)
		if err != nil {
			fmt.Println("read faild")
			break
		}

		msg, err := dp.Unpack(msgHead)

		msgData := make([]byte, msg.GetMsgLen())

		_, err = conn.Read(msgData)

		fmt.Printf("recv: %s, len = %d\n", msgData, cnt)

		time.Sleep(time.Second)
	}
}

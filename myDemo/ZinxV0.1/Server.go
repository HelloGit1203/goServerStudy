package main

import (
	"DAY03/zinx/ziface"
	"DAY03/zinx/znet"
	"fmt"
)

//ping test
type PingRouter struct {
	znet.BaseRouter
}

//test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle..")

	fmt.Println("recv from client :msgID =", request.GetMsgID(),
		", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

//test Handle
func (hr *HelloRouter) Handle(request ziface.IRequest) {
	//fmt.Println("Call Router Handle..")

	//fmt.Println("recv from client :msgID =", request.GetMsgID(),
	//", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("hello..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("--->DoConnectionBegin is Called...")
	conn.SendMsg(202, []byte("DoConnection Begin"))
}

func DoConnectionEnd(conn ziface.IConnection) {
	fmt.Println("--->DoConnectionEnd is Called...")
}

func main() {
	// create server
	s := znet.NewServer("[znix V0.1]")

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionEnd)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	//start server
	s.Serve()
}

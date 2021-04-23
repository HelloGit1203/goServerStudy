package znet

import (
	"fmt"
	"net"

	"DAY03/zinx/utils"
	"DAY03/zinx/ziface"
)

//iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//server name
	Name string
	//ip version
	IPVersion string
	//ip
	IP string
	//port
	Port int
	//server注册的连接对应的处理业务
	MsgHandler ziface.IMsgHandle

	//ConnMgr
	ConnMgr ziface.IConnManager

	OnConnStart func(ziface.IConnection)

	OnConnStop func(ziface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[Start]Server Listenner at IP:%s, Port:%d, is starting\n", s.IP, s.Port)
	go func() {
		s.MsgHandler.StartWorkerPool()

		//1. get tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}

		//2. listen
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen :", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx erver succ, ", s.Name, " succ, Listenning..")

		var cid uint32

		//3. accept
		for {

			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept error :", err)
				continue
			}

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("conn too many")
				conn.Close()
				//TODO
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	//TODO
	fmt.Println("Zinx server name = ", s.Name, " stop")
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO

	select {}
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add router success")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("----> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("----> Call OnConnStop")
		s.OnConnStop(conn)
	}
}

/*
inti server modle
*/

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}

	return s
}

package znet

import (
	"DAY03/zinx/utils"
	"DAY03/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClose bool

	ExitChan chan bool

	//该链接处理的方法router
	MsgHandler ziface.IMsgHandle

	msgChan chan []byte

	TcpServer ziface.IServer
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClose:    false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		TcpServer:  server,
	}

	c.TcpServer.GetConnMgr().Add(c)

	return c
}
func (c *Connection) startReader() {
	fmt.Println("reader goroutine is running...")

	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is ", c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConection(), headData)
		if err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}

		msg.SetData(data)

		//得到当前conn数据的Request数据
		req := Request{
			conn:    c,
			message: msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandle(&req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("Writer Goroution is running")

	defer fmt.Println(c.GetRemoteAddr().String(), "conn Writer exit")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID = ", c.ConnID)

	go c.startReader()

	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	if c.isClose {
		return
	}

	c.isClose = true

	c.TcpServer.CallOnConnStop(c)

	c.Conn.Close()

	c.ExitChan <- true

	c.TcpServer.GetConnMgr().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgid uint32, data []byte) error {

	if c.isClose {
		return errors.New("Connection closed when send msg")
	}

	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgid, data))
	if err != nil {
		return errors.New("package faild")

	}

	c.msgChan <- binaryMsg

	return nil
}

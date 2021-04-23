package znet

import "DAY03/zinx/ziface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection
	//客户端请求的数据
	message ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.message.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.message.GetMsgId()
}

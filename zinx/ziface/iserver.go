package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	// run server
	Serve()

	//路由功能，给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(uint32, IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(IConnection))

	SetOnConnStop(func(IConnection))

	CallOnConnStart(IConnection)

	CallOnConnStop(IConnection)
}

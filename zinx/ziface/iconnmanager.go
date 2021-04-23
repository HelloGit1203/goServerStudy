package ziface

type IConnManager interface {
	//添加连接
	Add(IConnection)
	//删除链接
	Remove(IConnection)

	Get(uint32) (IConnection, error)

	Len() int

	ClearConn()
}

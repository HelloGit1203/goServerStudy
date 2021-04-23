package ziface

/*

 */

type IMsgHandle interface {
	//执行对应的Router消息处理方式
	DoMsgHandle(request IRequest)
	//
	AddRouter(msgID uint32, router IRouter)

	StartWorkerPool()

	SendMsgToTaskQueue(IRequest)
}

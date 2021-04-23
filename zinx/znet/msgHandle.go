package znet

import (
	"DAY03/zinx/utils"
	"DAY03/zinx/ziface"
	"fmt"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter

	WorkerPoolSize uint32

	TaskQueue []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandle) DoMsgHandle(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), "is NOT FOUND!")
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeat api, msgID=" + fmt.Sprint(msgID))
	}

	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, "succ!")
}

func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, mh.WorkerPoolSize)

		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}

}

func (mh *MsgHandle) StartOneWorker(WorkerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", WorkerID, "is started...")

	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandle(request)
		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize

	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		"request MsgID = ", request.GetMsgID(), "to WorkerId = ", workerID)

	mh.TaskQueue[workerID] <- request
}

package znet

import "DAY03/zinx/ziface"

type BaseRouter struct {
}

//在处理conn业务之前的狗子方法
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

//处理业务主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

//处理业务之后的方法
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}

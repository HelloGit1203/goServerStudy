package ziface

type IRouter interface {
	//在处理conn业务之前的狗子方法
	PreHandle(request IRequest)

	//处理业务主方法
	Handle(request IRequest)

	//处理业务之后的方法
	PostHandle(request IRequest)
}

package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	// AddRouter 路由功能：给当前服务器注册一个路由方法，供客户端的连接调用
	AddRouter(router IRouter)
}

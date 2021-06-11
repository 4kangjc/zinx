package ziface

import "net"

type IConnection interface {
	Start()   									// 启动连接
	Stop()										// 停止连接
	GetTCPConnection() *net.TCPConn				// 获取当前连接的的socket
	GetConnID() uint32                          // 获取当前连接的ID
	RemoteAddr() net.Addr						// 获取远程客户端的地址
	Send(data []byte) error						// 发送数据给远程客户端
}

type HandleFunc func(*net.TCPConn, []byte, int) error
package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	conn *net.TCPConn					// 当前连接的socket
	connID uint32						// 连接的ID
	isClose bool						// 当前连接的状态
	ExitChan chan bool 					// 告知当前已经退出/停止的channel
	Router ziface.IRouter				// 该连接处理的方法
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection {
		conn: conn,
		connID: connID,
		Router: router,
		isClose: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// StartReader 读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.connID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf[:cnt],
		}
		// 从路由中找到注册绑定的Conn对应的router调用
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start()... connID = ", c.connID)
	go c.StartReader()

}

func (c *Connection) Stop() {
	fmt.Println("Connection Stop()... connID = ", c.connID)
	if c.isClose {
		return
	}
	c.isClose = true
	c.conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
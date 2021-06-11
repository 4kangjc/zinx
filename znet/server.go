package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name string
	IPversion string
	IP string
	Port int
	Router ziface.IRouter
}


func (s *Server) Start() {
	go func() {
		fmt.Printf("[Start] Server Listener at IP : %s, Port %d, is starting\n", s.IP, s.Port)
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		listener, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPversion, " err ", err)
			return
		}
		fmt.Println("start Zinx server successful ", s.Name, "Listening...")
		var cid uint32
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			//go dealConn.Start()
			dealConn.Start()

		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Server() {
	s.Start()
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Successful!!!")
}

func NewServer(name string) ziface.IServer {
	s := &Server {
		Name: name,
		IPversion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
		Router: nil,
	}
	return s
}
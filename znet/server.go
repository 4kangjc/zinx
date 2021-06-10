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
		fmt.Printf("start Zinx server successful ", s.Name, "Listenning...")
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			go func() {
				buf := make([]byte, 512)
				for {
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("recv buf is %s\n", buf)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Server() {
	s.Start()
	select {}
}


func NewServer(name string) ziface.IServer {
	s := &Server {
		Name: name,
		IPversion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}
package server

import (
	"fmt"
	"net"
)



type Server struct{
	address string
	listner net.Listener
}

func NewServer(address string) *Server{
	listner, _ := net.Listen("tcp", address)
	return &Server{
		address: address,
		listner: listner,
	}
}

func (s *Server) Start(){
	for {
		conn, _ := s.listner.Accept()
		fmt.Println(conn)
	}
}
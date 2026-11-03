package server

import (
	"fmt"
	"net"
	"redis-lite/internal/client"

	"redis-lite/internal/parser"
	"redis-lite/internal/framer"
	"redis-lite/internal/protocol" )



type Server struct{
	Address string
	Listner net.Listener
	Clients map[*client.Client]struct{}
}

func NewServer(address string) *Server{
	listner, _ := net.Listen("tcp", address)
	return &Server{
		Address: address,
		Listner: listner,
		Clients: make(map[*client.Client]struct{}),
	}
}

func (s *Server) Start(){
	for {
		conn, _ := s.Listner.Accept()
		newClient := client.NewClient(conn)
		s.AddClientToGlobalList(newClient)
		go s.HandleClientConnection(newClient.Conn)

	}
}


func (s *Server) HandleClientConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	clientframer := framer.Newframer()
	for {
		n , err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)

		}
		validmessages := clientframer.Feed(buffer[:n])
		//clear the buffer for safety
		clear(buffer)
		commands := []*protocol.COMMAND{}
		for _ , message :=  range validmessages{
			command, err := parser.Parser(message)
			if err != nil {
				break
				//if [valid, valid, invalid, valid] -> last valid wont be processed
			}
			commands = append(commands, command)
		}

	}
	


}
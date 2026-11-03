package server

import (
	"redis-lite/internal/client"
)


func (s *Server) AddClientToGlobalList(client *client.Client){
	s.Clients[client] = struct{}{}
}
package main

import (
	"redis-lite/internal/server"
)
	

func main(){

	server := server.NewServer(":3000")
	server.Start()

}
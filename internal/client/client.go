package client

import (
	"net"
)


type Client struct{
	Conn net.Conn
	ClientID string
}


func NewClient(conn net.Conn) *Client{
	return &Client{	
		Conn: conn,
	}
}
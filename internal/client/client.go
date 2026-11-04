package client

import (
	"net"
	s "redis-lite/internal/store"
)


type Client struct{
	Conn net.Conn
	ClientID string
	CurrentStore *s.Store
}


func NewClient(conn net.Conn, store *s.Store) *Client{
	return &Client{	
		Conn: conn,
		CurrentStore: store,
	}
}

func (c *Client) SendDataToClient(data string) error {
	//data could be send a bit so have a re-try logic

	databyte := []byte(data)
	databyte = append(databyte, byte('\n'))
	_, err := c.Conn.Write(databyte)
	if err != nil {
		return err
	}
	return nil
}
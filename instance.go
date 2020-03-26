package main

import (
	"fmt"
	"log"
	"net"
)

type SocketProcess interface {
	Listen(s *Server) (net.Conn, error)
	OnStartError(err error)
	Read(conn net.Conn) (string, interface{}, error)
	Write(conn net.Conn)
}

func Process(s *Server, sp SocketProcess) {
	conn, err := sp.Listen(s)
	if err != nil {
		log.Printf("ID: %d TYPE: %s: ERROR: %s", s.id, s.name, err)
		sp.OnStartError(err)
	}
	client := Client{id, &conn}
	append(clients, &client)
	id = id + 1
	fmt.Println("%s服务器端启动:  %s", s.name, addr.String())
	defer s.OnClose(&client)
	defer conn.Close()
	for {
		mt, message, err := sp.Read(conn)
		if err != nil {
			s.OnError(&client)
		}
		s.OnMessage(&client, message)
	}
}

type Onstart func()
type OnOpen func()
type OnConnect func(cli *Client)
type OnMessage func(cli *Client, msg interface{})
type OnError func(cli *Client)
type OnClose func(cli *Client)

var (
	id      int64 = 0
	clients []*Client
)

type Client struct {
	id   int64
	conn *net.Conn
}

type Server struct {
	id        int64
	name      string
	addr      net.Addr
	OnMessage OnMessage
	OnError   OnError
	OnClose   OnClose
}

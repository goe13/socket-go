package main

import (
	"fmt"
	"go/types"
	"log"
	"net"
)

type socket_type uint8

const (
	UDP        socket_type = 0
	TCP        socket_type = 1
	WEB_SOCKET socket_type = 2
)

type SocketProcess interface {
	Listen(*Server) (*Conn, error)
	OnStartError(error)
	Read(net.Conn) (string, interface{}, error)
	Write(net.Conn)
}

type ConnRW interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Conn struct {
	conn interface{}
}

func Process(s *Server) {
	w := &Worker{s.sType, run}
	for {
		conn, err := w.Listen(s.addr)
		if err != nil {
			log.Printf("ID: %d TYPE: %s: ERROR: %s", s.id, s.name, err)
		}
		client := &Client{id, conn}
		clients = append(clients, client)
		id = id + 1
		fmt.Println("%s服务器端启动:  %s", s.name, s.addr)
		defer s.OnClose(client)
		defer conn.Close()
		for {
			message, err := conn.Read()
			if err != nil {
				s.OnError(client, err)
			}
			s.OnMessage(client, message)
		}
	}
}

type Worker struct {
	sType socket_type
	run   func()
}

var (
	id      int64 = 0
	clients []*Client
	servers []*Server
)

type Client struct {
	id   int64
	conn *Conn
}

type Server struct {
	id        int64
	sType     socket_type
	addr      string
	OnMessage func(*Client, []byte)
	OnError   func(error)
	OnClose   func(*Client)
	OnConnect func(*Client)
	OnStart   func()
	OnOpen    func()
}

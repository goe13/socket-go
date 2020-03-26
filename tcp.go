package main

import (
	"fmt"
	"net"
	"time"
)

type TCPWorker struct {
}

func (w *TCPWorker) Process(s *Server) {
	addr, err := net.ResolveTCPAddr("tcp", s.addr)
	s.OnError(err)
	tcpListen, err := net.ListenTCP("tcp", addr)
	s.OnError(err)
	fmt.Println("tcp服务器端启动:  ", addr.String())
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			continue
		}
		c := &Conn{}
		client := &Client{id, c}
		clients = append(clients, client)
		var buf [1024]byte
		for {
			readSize, err := conn.Read(buf[0:])
			s.OnError(err)
			remoteAddr := conn.RemoteAddr()
			fmt.Println("来自远程ip:", remoteAddr, " 的消息:", string(buf[0:readSize]))
			_, err2 := conn.Write([]byte(string(buf[0:readSize]) + " " + time.Now().String()))
			//一定要执行下面的return 才能监听到客户端主动断开，服务器端对本次conn进行close处理 dealErrorWithReturn不能达到这个效果。
			if err2 != nil {
				fmt.Println("当前用户 ", conn.RemoteAddr(), " 主动断开链接！")
				conn.Close() //正常链接情况下，handlerConn不会释放出来到这里 当客户端强制断开，才会return到这里，吧当前conn关闭
				s.OnClose(client)
				return
			}
		}
	}
}

func (c *Conn) Read() ([]byte, error) {
	var buf [1024]byte
	readSize, err := c.conn.Read(buf[0:])
	if err != nil {
		return buf[0:], err
	}
	return buf[0:readSize], nil
}
func (c *Conn) Write(b []byte) error {
	_, err := c.conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

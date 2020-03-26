package main

import (
	"fmt"
	"net"
)

func tcp() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:18880")
	dealErrorWithExist(err)
	tcpListen, err := net.ListenTCP("tcp", addr)
	dealErrorWithExist(err)
	fmt.Println("tcp服务器端启动:  ", addr.String())
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			continue
		}
		handlerConn(conn)
		conn.Close() //正常链接情况下，handlerConn不会释放出来到这里 当客户端强制断开，才会return到这里，吧当前conn关闭
		fmt.Println("当前用户 ", conn.RemoteAddr(), " 主动断开链接！")
	}
}

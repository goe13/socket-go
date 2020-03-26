package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func dealErrorWithExist(err error) {
	//有异常会停止进程
	if err != nil {
		log.Fatal("运行异常", err)
	}
}
func dealErrorWithReturn(err error) {
	//有异常会停止进程
	if err != nil {
		return
	}
}

func handlerConn(conn net.Conn) {
	var buf [1024]byte
	for {
		readSize, err := conn.Read(buf[0:])
		dealErrorWithReturn(err)
		remoteAddr := conn.RemoteAddr()
		fmt.Println("来自远程ip:", remoteAddr, " 的消息:", string(buf[0:readSize]))
		_, err2 := conn.Write([]byte(string(buf[0:readSize]) + " " + time.Now().String()))
		//一定要执行下面的return 才能监听到客户端主动断开，服务器端对本次conn进行close处理 dealErrorWithReturn不能达到这个效果。
		if err2 != nil {
			return
		}
	}
}

func main() {
	go ws()
	fmt.Println("ws is run")
	tcp()
	fmt.Println("tcp is run")
}

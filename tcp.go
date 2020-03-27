package socketgo

import (
	"fmt"
	"net"
)

type TCPWorker struct {
	s *Server
}
type TCPConn struct {
	mConn net.Conn
}

func (c *TCPConn) Write(b []byte) error {
	_, err := c.mConn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (w *TCPWorker) Run() {
	addr, err := net.ResolveTCPAddr("tcp", w.s.addr)
	w.s.OnError(err)
	tcpListen, err := net.ListenTCP("tcp", addr)
	w.s.OnError(err)
	fmt.Println("tcp服务器端启动:  ", addr.String())
	ch := make(chan bool, 1000)
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		c := &TCPConn{conn}
		counter++
		connector := &Connector{counter, c}
		connectors = append(connectors, connector)
		w.s.OnConnect(connector)
		ch <- true
		go func() {
			defer func() {
				<-ch
			}()
			defer w.s.OnClose(connector)
			defer conn.Close()
			for {
				var buf [1024]byte
				readSize, err := conn.Read(buf[0:])
				if err != nil {
					w.s.OnError(err)
					break
				}
				if readSize > 0 {
					w.s.OnMessage(connector, buf[0:readSize])
				}
				w.s.OnError(err)
				//remoteAddr := conn.RemoteAddr()
			}
		}()
	}

}

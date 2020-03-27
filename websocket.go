package socketgo

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSWorker struct {
	s *Server
}
type WSConn struct {
	mConn *websocket.Conn
}

func (c *WSConn) Write(b []byte) error {
	err := c.mConn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}
	return nil
}

func (w *WSWorker) Run() {
	upgrader := websocket.Upgrader{
		//取消跨域限制
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		c := &WSConn{conn}
		counter++
		connector := &Connector{counter, c}
		connectors = append(connectors, connector)
		w.s.OnConnect(connector)
		defer w.s.OnClose(connector)
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				w.s.OnError(err)
				break
			}
			w.s.OnMessage(connector, []byte(message))
		}
	})
	fmt.Println("websocket服务器端启动:  ", w.s.ADDR)
	log.Fatal(http.ListenAndServe(w.s.ADDR, nil))
}

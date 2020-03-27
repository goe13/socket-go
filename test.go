package socketgo

import "fmt"

func main() {

	serv := &Server{
		sType: TCP,
		addr:  "127.0.0.1:18880",
		OnMessage: func(conn *Connector, b []byte) {
			fmt.Println(string(b))
			conn.conn.Write([]byte("tcp connect \n"))
			//sendToClient(conn.id,[]byte("tcp"))
			//panic(11)
		},
		OnError: func(err error) {

		},
		OnClose: func(conn *Connector) {
			fmt.Println("closed")
		},
		OnConnect: func(conn *Connector) {

		},
		OnStart: func() {

		},
		OnOpen: func() {

		},
	}
	tp := &Pr{[]Worker{}}
	tp.AddServer(serv)
	tp.RunAll()
}

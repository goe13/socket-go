package socketgo

import "fmt"

func main() {

	serv := &Server{
		S_TYPE: TCP,
		ADDR:   "127.0.0.1:18880",
		OnMessage: func(conn *Connector, b []byte) {
			fmt.Println(string(b))
			conn.conn.Write([]byte("tcp connect \n"))
			//SendToClient(conn.ID,[]byte("tcp"))
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
	pr := GetProcessor(serv)
	pr.OnStart = func() {
		fmt.Println("start running !")
	}
	pr.RunAll()
}

package socketgo

import "fmt"

func main() {

	serv := &Server{
		S_TYPE: TCP,
		ADDR:   "127.0.0.1:18880",
		OnMessage: func(conn *Connector, b []byte) {
			fmt.Println(string(b))
			//conn.conn.Write([]byte("tcp connect \n"))
			//SendToClient(conn.ID,[]byte("tcp"))
			//panic(11)
		},
	}
	pr := GetProcessor(serv)
	pr.OnStart = func() {
		fmt.Println("start running !")
	}
	pr.RunAll()
}

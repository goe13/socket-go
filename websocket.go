package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "127.0.0.1:18881", "http service address")

var upgrader = websocket.Upgrader{
	//取消跨域限制
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func root(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func ws() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

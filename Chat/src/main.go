package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"./message"
)

var clients = make(map[*websocket.Conn]bool)	//客户端连接
var broadcast = make(chan message.Message)	//消息管道

var upgrader = websocket.Upgrader{}

func HandlerConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg message.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		broadcast <-msg
	}
}

func handlerMessage() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("../public"))	//设置http的主目录
	http.Handle("/", fs)	//将根目录与public目录绑定 当用户访问/时 默认访问服务器的../public目录

	go handlerMessage()

	log.Println("http server started on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

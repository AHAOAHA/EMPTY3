package main
//webSocket是HTML5开始提供的浏览器和服务器之间进行全双工通讯的网络技术。

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)




func main() {
	err := http.ListenAndServe("0.0.0.0", Headler)
	if err !=nil {
		log.Fatal(err)
	}
}

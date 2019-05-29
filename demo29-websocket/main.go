package main
//使用websocket协议

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)



func main() {
	http.Handle("/websocket", websocket.Handler(Echo))

	http.HandleFunc("/web", web)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	//http.ListenAndServe:
	//1.实例化Server对象,会创建服务端的套接字
	//2.调用Server的ListenAndServe()
	//3.调用net.Listen("tcp", addr)监听端口
	//4.启动for循环，接收连接请求
	//5.当有请求到达时，开启一个goroutine处理请求
	//6.
}

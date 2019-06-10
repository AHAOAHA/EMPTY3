package main

import (
	"flag"
	"fmt"
	"net"
)

//ECHO服务器，客户端代码

func main() {
	flag.PrintDefaults()
	flag.Parse()
	Con, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer Con.Close()

	buffer := make([]byte, 1024)

	for {
		var req string
		fmt.Printf("Client# ")
		_, _ = fmt.Scanln(&req)

		fmt.Println("req:", req, "len:", len(req))
		_, err = Con.Write([]byte(req))
		fmt.Println("Write not stop")
		if err != nil {
			return
		}

		n, _ := Con.Read(buffer)
		rsp := buffer[:n]
		fmt.Println("Serv:", string(rsp))
	}
}

package main

import (
	"fmt"
	"net"
)

// 使用tcp连接构建简单的echo服务器

var dict = make(map[string]string, 10)


func main() {
	dict["ahaoo"] = "张昊"
	dict["xixi"] = "嘻嘻"
	// 创建tcp连接 listenr
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	defer func(){
		err := ln.Close()
		if err != nil {
			panic(err)
		}
	}()

	for{
		cli, err := ln.Accept()
		if err != nil {
			continue
		}

		go handler(cli)
	}
}


// 处理函数
func handler(cli net.Conn) {
	defer func(){
		err := cli.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Get a Cli")
	buffer := make([]byte, 1024)

	for {
		n, err := cli.Read(buffer)
		if err != nil {
			return
		}

		req := buffer[:n]
		fmt.Println("Cli:", string(buffer))

		rsp, ok := dict[string(req)]
		if !ok {
			_, _ = cli.Write([]byte("have no mem"))
		}
		_, err = cli.Write([]byte(rsp))

		if err != nil {
			return
		}
	}
}
package main

import (
	"context"
	"log"
	"time"
	"os"
	"fmt"
)

var logg *log.Logger

func someHandler() {
	//创建出一个context,同时返回它的取消函数,同时将k,v绑定到该context上
	ctx := context.WithValue(context.Background(), k, v)
	//使用刚刚创建的context管理doStuff协程
	go DoValueStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func DoValueStuff(ctx context) {
	for {
		time.Sleep(time.Second)
		if v, ok := ctx.Value(3).(); ok {
			logg.Printf("value set")
			fmt.Println(v)
		} else {
			logg.Printf("value not set")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

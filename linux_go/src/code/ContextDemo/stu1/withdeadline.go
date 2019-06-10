package main

import (
	"context"
	"log"
	"time"
	"os"
)

var logg *log.Logger

func someHandler() {
	//创建出一个context,同时返回它的取消函数,同时该context包含超时时间timeout
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	//使用刚刚创建的context管理doStuff协程
	go doStuff(ctx)

	//在context到达超时时间的时候，context就已经被取消

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

//每一秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		//每次循环等待1s
		time.Sleep(time.Second)

		//在前10s会一直执行default
		select {
			//监听context是否取消，如果context已经取消，则该协程没有继续工作下去的必要
			case <- ctx.Done():
				logg.Printf("done")
				return
			default:
				logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

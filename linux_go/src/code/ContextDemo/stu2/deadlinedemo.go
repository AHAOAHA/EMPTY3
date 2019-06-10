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
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	//使用刚刚创建的context管理doStuff协程
	//go doStuff(ctx)
	go DoTimeOutStuff(ctx)

	//在context到达超时时间的时候，context就已经被取消

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func DoTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			logg.Printf("deadline set")
			if time.Now().After(deadline) {
				logg.Printf(ctx.Err().Error())
				return
			}
		} else {
			logg.Printf("deadline not set")
		}


		select {
			case <-ctx.Done():
				logg.Printf("done")
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

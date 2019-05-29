package main
//定时任务

import (
	"fmt"
	"time"
)

func NewTimerTestFunc() {
	t := time.NewTimer(time.Second * 3)


	 <-t.C	//会阻塞在此处，等待定时器触发
	 fmt.Println("NewTimer Trigger")
}

func AfterFuncTestFunc() {
	//AfterFunc本身不会阻塞
	time.AfterFunc(time.Second*3, func(){fmt.Println("AfterFuncTestFunc Trgger")})
}

func TickerTestFunc() {
	t := time.NewTicker(time.Second * 3)

	for {
		<-t.C
		fmt.Println("TickerFunc trgger")
	}
}

func main() {
	//NewTimerTestFunc()
	//AfterFuncTestFunc()
	//TickerTestFunc()
	fmt.Printf("%x%+x", 1, " haha")

}

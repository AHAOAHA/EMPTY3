package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go func(c chan int) {
		//go routine
		fmt.Println("this is routine 1")
		time.Sleep(time.Second * 3)
		c <- 1	//routine执行完成后向chan中写入1
	}(c)

	go func(c chan int) {
		<-c	//从chan中读数据，只有读出数据时间才会向下执行
		fmt.Println("this is routine 2")
	}(c)

	time.Sleep(time.Second * 60)
}

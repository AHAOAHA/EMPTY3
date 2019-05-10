//使用go run main.go只执行Once2
//使用go build + ./main 只执行Once1


package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var o sync.Once
	go func() {
		o.Do(Once1)
		fmt.Println("routine 1 run ...")
	}()

	time.Sleep(time.Second)

	go func() {
		o.Do(Once2)
		fmt.Println("routine 2 run ...")
	}()

	time.Sleep(time.Second * 10)
}

func Once1() {
	fmt.Println("hello")
}

func Once2() {
	fmt.Println("world")
}


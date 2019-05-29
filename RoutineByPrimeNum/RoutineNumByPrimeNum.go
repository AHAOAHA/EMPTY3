//测试机器可以同时创建出多少协程，可能会导致机器卡顿
package main

import (
	"fmt"
	"time"
)

func RoutineNum(sum uint64) {
	sum += 1
	fmt.Println("Routine num:", sum)
	go RoutineNum(sum)
	for {
		time.Sleep(time.Minute)
	}
}

func main() {
	var sum uint64 = 0
	go RoutineNum(sum)
	for {
		time.Sleep(time.Minute)
	}
}
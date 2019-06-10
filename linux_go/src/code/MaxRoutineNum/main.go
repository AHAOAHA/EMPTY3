package main

import (
	"fmt"
	"time"
)

func RoutineNum(num uint64) {
	num += 1
	fmt.Println("Routine num:", num)
	go RoutineNum(num)
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

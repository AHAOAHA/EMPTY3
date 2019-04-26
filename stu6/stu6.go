package main

import (
	"fmt"
)

func hello() {
	fmt.Println("this is goroutine")
}

func main() {
	go hello()
	fmt.Println("this is main")
	for {

	}
}

package main

import (
	"fmt"
)

//golang Select不会循环检测
func TestRound() {
	select {
	default:
		fmt.Println("default Handler")
	}
}

func main() {
	TestRound()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("routine 1")
		for {
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("routine 2")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("done ...")
}

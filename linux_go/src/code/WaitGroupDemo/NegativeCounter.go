package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("routine 1")
		wg.Done()
		wg.Done()
		wg.Done()
	}()

	go func() {
		fmt.Println("routine 2")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("done ...")
}

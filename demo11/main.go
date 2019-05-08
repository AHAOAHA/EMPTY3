package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func main() {
	err := errors.New("this is error")
	fmt.Println(err)
}



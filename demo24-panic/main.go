package main

import (
	//"fmt"
	"fmt"
)

func RecoverPanicTestFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("painc has been recover")
		}
	}()

	//panic(1)	//当前函数内部的panic可以被recover
	PanicTestFunc()	//内部调用的函数也可以被panic
}

func PanicTestFunc() {
	panic(1)
}

func main() {
	//PanicTestFunc()
	RecoverPanicTestFunc()
}

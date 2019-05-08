package main

import (
	"fmt"
)

//定义变长参数列表
func MyPrint(a string, who ...int) {
	for _, v := range who {
		fmt.Printf("%d ", v)
	}
}



func test() int {
	return 4567
}

type Func interface {
	Test()
}
func main() {
	var f Func
	f.Test()
}

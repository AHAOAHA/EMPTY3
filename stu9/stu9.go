package main

//#include <stdio.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
)

//定义变长参数列表
func MyPrint(a string, who ...int) {
	for _, v := range who {
		fmt.Printf("%d ", v)
	}
}

func main() {
	var a int = 5
	MyPrint("", 1, 2, 3, 4, 5, 6, 7)
	C.printf("%d\n", a)
}

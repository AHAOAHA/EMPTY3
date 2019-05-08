//有返回值的协程

//有返回值的协程，返回值会被直接丢弃
package main

import (
	"fmt"
)

func RetValFunc() int {
	return 0
}

func main() {
	var ret int
	go RetValFunc()
	fmt.Println(ret)
}

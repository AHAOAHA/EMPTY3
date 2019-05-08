//init test
package main

//main函数在init函数执行完成之后才会开始执行
import (
	"fmt"
	"time"
)

func init() {
	for {
		time.Sleep(time.Second)
	}
}
func main() {
	fmt.Println("hello world!")
}

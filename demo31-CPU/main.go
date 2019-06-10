// CPU密集型的协程会一直被运行，不会让出CPU给其他协程，编写代码测试之

package main

import (
	"fmt"
	"time"
)

//一个CPU密集型的协程
func CPUIntensive() {
	var i uint64
	for {
		i++
	}
}

//IO密集
func Echo() {
	for{
		fmt.Println("hello world!")
		time.Sleep(time.Second)
	}
}

func main() {
	go CPUIntensive()
	go Echo()
	time.Sleep(time.Second)
}

//在单核机器上进行测试时，如果CPU密集型协程在主执行流退出前启动，则程序会永久阻塞，并且Echo协程不会被调度

package main

import (
	"fmt"
	"sync"
)

func main() {
	//建立Pool对象
	var pool = &sync.Pool {New: func() interface{} {return "hello Newer"}}

	var tmp string = "ahaoozhang"
	pool.Put(tmp)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())

	var str interface{} = "1122331"
	fmt.Println(str)
}

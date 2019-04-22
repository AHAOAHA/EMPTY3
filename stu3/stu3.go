package main

import "fmt"

func test() (a, b, c int) /*第二个括号中代表返回值*/ {
	return 1, 2, 3
}

func main() {
	//Print和Println的区别
	a := 10
	//一段一段处理 自动添加换行 变量之前自动添加空格 类似于cpp的cout
	fmt.Println("a =", a)
	//类似于C语言中的printf
	fmt.Printf("a = %d\n", a)

	//调用函数
	var b, c int     //用来接收函数返回值
	b, _, c = test() //使用匿名变量可以丢弃掉多余的变量
	fmt.Println("b =", b, "c =", c)
}

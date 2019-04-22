package main

import "fmt"

// fmt包含IO相关函数

//go语言关键字
//int default break func interface select case defer go map struct
//chan else goto package switch const fallthrough if range type
//continue for import return var

//内建常量
//true false iota nil

//内建类型
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64 uintptr
//float32 float64 complex128 complex64
//bool byte rune string error

//内建函数
//make len cap new append copy close delete
//complex real imag
//panic recover

//变量的使用
func main() {
	//1. 变量的声明格式：var [变量名] [变量类型]
	//变量声明之后必须进行使用，否则就会报错
	//只有声明没有初始化的变量的默认值为0
	var a int
	var b, c int
	fmt.Println("a = ", a)
	fmt.Println("b = ", b, "c = ", c)

	//变量初始化
	var d int = 10
	fmt.Println("d = ", d)
	//var e = 10 int 错误方式
	//fmt.Println("e = ", e)

	//变量赋值
	a = 5
	fmt.Println("a = ", a)

	//自动推导类型 通过初始化的值的类型推导变量类型 :=
	//:= 先声明变量的类型 再给变量赋值
	//对于同一个变量 := 只能使用一次 这可以作为 := 和 = 的区别
	f := 5
	fmt.Printf("f type is %T\n", f) //Printf无自动换行 %T代表变量的类型
	fmt.Println("f = ", f)

	//多重推导初始化
	g, h := 3, 4
	fmt.Printf("g = %d, h = %d\n", g, h)

	//交换两个变量的值
	g, h = h, g
	fmt.Printf("g = %d, h = %d\n", g, h)

	//匿名变量
	var tmp int
	tmp, _ = h, g // _ 代表匿名变量 此处g值会被丢弃
	fmt.Printf("tmp = %d\n", tmp)

}

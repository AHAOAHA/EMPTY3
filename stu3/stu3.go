package main

import (
	"fmt"
	"time"
)

func test() (a, b, c int) /*第二个括号中代表返回值*/ {
	return 1, 2, 3
}

func main() {
	//Print和Println的区别
	a := 1
	//一段一段处理 自动添加换行 变量之前自动添加空格 类似于cpp的cout
	fmt.Println("a =", a)
	//类似于C语言中的printf
	fmt.Printf("a = %d\n", a)

	//调用函数
	var b, c int     //用来接收函数返回值
	b, _, c = test() //使用匿名变量可以丢弃掉多余的变量
	fmt.Println("b =", b, "c =", c)

	//常量的使用
	//	变量声明关键字：var
	//	常量声明关键字： const
	const d int = 10
	const e = 20 //常量的推导不能使用 :=
	fmt.Println("d =", d)
	fmt.Println("e =", e)

	//多个变量或常量定义
	f, g := 20, 1.878
	fmt.Println("f =", f, "g =", g)

	var (
		i int
		j float64
	)
	i = 3
	j = 3.14
	fmt.Println("i =", i, "j =", j)

	var (
		o = 1
		p = 3.14
	//q := 2 使用war声明变量时不能使用:=进行推导 可以直接使用=进行推导
	)
	fmt.Println("o =", o, "p =", p)

	const (
		k int     = 5
		l float64 = 3.654
	)
	fmt.Println("k =", k, "l =", l)

	const (
		m = 10
		n = 3.145926
	)
	fmt.Println("m =", m, "n =", n)

	//枚举
	const (
		//iota为常量自动生成器， 每隔一行自动累加1
		//iota给常量赋值使用
		r = iota
		s = iota
		t = iota
	)	
	fmt.Printf("r = %d, s = %d, t = %d\n", r, s, t)
	const (
		//iota遇到const就会重置为0
		u = iota
		v = iota
	)
	fmt.Printf("u = %d, v = %d\n", u, v)

	//在同一个const中，可以只写一个iota
	const (
		r1 = iota
		s1
		t1
	)
	fmt.Printf("r1 = %d, s1 = %d, t1 = %d\n", r1, s1, t1)

	//如果是同一行，值都一样
	const (
		u1         = iota
		v1, v2, v3 = iota, iota, iota
	)
	fmt.Printf("u1 = %d, v1 = %d, v2 = %d, v3 = %d\n", u1, v1, v2, v3)

	//基础数据类型
	//类型 名称     大小  零值  说明
	//bool 布尔类型 1字节 false 其数字不为真即为假，不能用数字代表true or false
	//byte 字节型   1字节  0    uint8别名
	//rune 字符类型 4字节  0    专门用于存储unicode编码，等价于uint32
	//int,uint 整型 4字节/8字节 0 32位或64位
	//int8,uint8 整型 1字节 0 -128~127，0~255
	//int16,uint16 整型 2字节 0 -32768~32767，0~65535
	//int32,uint32 整型 4字节 0 -21亿~21亿，0~42亿
	//int64,uint64 整型 8字节 0
	//float32 浮点型 4字节 0.0 小数精确到7位
	//float64 浮点型 8字节 0.0 小数精确到15位
	//complex64 复数类型 8
	//complex128 复数类型 16
	//uintptr 整型 4字节或8字节 0 足以存储指针的uint32或uint64整数
	//string 字符串               utf-8字符串

	//复数声明
	var z complex64
	z = 1 + 3i
	fmt.Println("z =", z)
	//通过内建函数取实部/虚部
	fmt.Println("real =", real(z), "imag =", imag(z))

	//输入
	var input int
	//fmt.Scanf("%d", &input) //阻塞等待用户输入
	//fmt.Scan(&input)
	fmt.Println("inoput =", input)

	//类型转换
	flag := true
	//fmt.Printf("flag = %d\n", flag)
	fmt.Printf("flag = %t\n", flag)
	//bool类型不能转换为整型 整型也不能转换为bool类型

	//这种不能转换的类型叫做不兼容类型

	var ch byte
	ch = 'a'
	fmt.Printf("ch = %c\n", ch)
	var _ch int
	_ch = int(ch)
	fmt.Println("_ch =", _ch)

	//类型别名 类似于typedef
	//type [别名] [原名]
	type bigint int64
	type (
		smallint  int8
		usmallint uint8
	)
	var test bigint = 1
	fmt.Println("test =", test)

	//if语句
	str := "ahaoaha"
	if str == "ahaoaha" { //左括号和if必须在同一行
		fmt.Println("ahaoaha Ok!")
	}
	//if支持一个初始化语句 ;之后为判断条件 ;之前为初始化语句
	if str1 := "ahao"; str == "ahao" {
		fmt.Println(str1, "Ok!")
	}

	//else / else if必须和之前的右括号以及自己的左括号在同一行
	a1 := 10
	if a1 == 10 {
		fmt.Println("a1 == 10")
	} else {
		fmt.Println("a1 != 10")
	}

	//switch语句
	num := 1
	switch num { //go语言保留break关键字 跳出switch循环 不写默认添加
	case 1:
		fmt.Println("num == 1")
		break
	case 2:
		fmt.Println("num == 2")
		break
	case 3:
		fmt.Println("num == 3")
		break
	case 4:
		fmt.Println("num == 4")
		break
	default:
		fmt.Println("num == x")
	}

	//更加灵活的switch
	//支持一个初始化语句
	switch num1 := 1; num1 {
	case 1:
		fmt.Println("num1 == 1")
	case 2:
		fmt.Println("num1 == 2")
	case 3, 4, 5: //case后面可以跟多个条件
		fmt.Println("num1 == 3")
	default:
		fmt.Println("num1 == x")
	}

	//可以没有条件
	score := 87
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 85:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	case score > 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	//for循环
	//格式：for 初始条件; 判断条件; 条件变换 {
	//TODO
	//}

	//实现1+2+3......+100
	var sum int
	for tmp := 1; tmp <= 100; tmp++ {
		sum += tmp
	}
	fmt.Println("sum =", sum)

	//range的使用
	str3 := "ahaoaha"
	//通过for打印字符
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c ", str3[i])
	}
	fmt.Printf("\n")
	//更好的写法 range 迭代打印每个元素
	//默认返回两个值，一个是元素位置，一个是元素本身
	for i, data := range str3 {
		fmt.Printf("str3[%d] = %c\n", i, data)
	}
	for i := range str3 { //此时第二个参数默认丢弃
		fmt.Printf("str3[%d] = %c\n", i, str3[i])
	}

	//break和continue
	//break可用于for\switch\select，而continue只能用于for循环
	//break跳出当前循环
	//continue跳出本次循环
	i = 0
	for { //for后面不加任何东西，这个循环条件永远为真
		time.Sleep(time.Second)
		i++
		if i == 5 {
			//break
			continue
		}
		fmt.Printf("i = %d\n", i)
	}

	//goto语句
	//用goto跳转必须在当前函数内定义标签
	//goto可以在任何地方，但是不能跨函数使用

}

package main

import (
	"fmt"
)

//函数
//函数定义的位置无特定要求

//函数定义格式
//func FumcName([参数列表])(o1 [返回值类型], o2 [返回值类型]) {
//函数体
//return ...
//}

//函数定义说明
//func :函数由关键字func开始声明
//FuncName :函数名，根据约定，函数名首字母小写即为private，首字母大写即为public
//参数列表 :函数可以有0个或者多个参数，参数格式为：[变量名] [类型]，如果有多个参数通过逗号分隔，不支持默认参数
//返回类型：
//1） 上面返回值声明了两个变量o1和o2(命名返回参数)，这个不是必须的，可以只有类型没有变量名
//2） 如果只有一个返回值且不声明返回值变量，那么可以省略，包括返回值的括号
//3） 如果没有返回值就直接省略最后的返回值信息
//4） 如果有返回值，那么必须在函数的内部添加return语句

//无参无返回值
func MyFuncFirst() {
	a := 1998
	fmt.Println("born in", a)
}

//有参无返回值,普通参数列表
func MyFuncSecond(a int) {
	//参数列表中声明变量无需var关键字
	a = 789
	fmt.Println("func a =", a)
}
func main() {
	MyFuncFirst()
	var a int = 10
	MyFuncSecond(a)
}

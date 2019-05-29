//空接口
//每一种类型都能匹配到空接口，也就是说可以使用空接口的作为函数的参数来接收任何类型的对象

package main

import "fmt"

type I interface {
	Get() string
}

func GetName(e interface{}) string {
	return e.(I).Get()
}

type Class struct {
	name   string
	gender string
}

func (c Class) Get() string {
	return c.name
}
func main() {
	c := Class{
		name:   "ahaoozhang",
		gender: "man",
	}
	fmt.Println(GetName(c))
}

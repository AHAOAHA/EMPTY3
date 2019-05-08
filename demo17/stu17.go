//接口类型判断demo

//在Go中，要判断传递给接口的数据类型，可以使用type switch得到 (type)只能在switch中使用
package main

import (
	"fmt"
	"time"
)

type GetStuInfo interface { //定义接口
	Name()
	Date()
}

type Class1 struct {
	name string
	date time.Time
}

func (s Class1) Name() {
	fmt.Println("class1 student name:", s.name)
}

func (s Class1) Date() {
	fmt.Println("class1 student date:", s.date)
}

type Class2 struct {
	name string
	date time.Time
}

func (s Class2) Name() {
	fmt.Println("class2 student name:", s.name)
}

func (s Class2) Date() {
	fmt.Println("class2 date:", s.date)
}

func IntFunc(p GetStuInfo) {
	switch p.(type) {
	case Class1:
		fmt.Println("Class1值传递")
	case *Class1:
		fmt.Println("*Class1引用传递")
	case Class2:
		fmt.Println("Class2值传递")
	case *Class2:
		fmt.Println("*Class引用传递")
	}

	p.Name()
	p.Date()
}

func main() {
	c1 := Class1{
		name: "ahaoozhang",
		date: time.Now(),
	}
	IntFunc(c1)
}

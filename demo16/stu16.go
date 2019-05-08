//interface

package main

import (
	"fmt"
	"time"
)

type Student struct { //结构体
	name   string
	date   time.Time
	gender string
}

func (s Student) Name() string {
	fmt.Println("name:", s.name)
	return s.name
}

func (s Student) Gender() string {
	fmt.Println("gender:", s.gender)
	return s.gender
}

func (s Student) Date() string {
	fmt.Println("date:", s.date)
	return s.name
}

type PrintUserInfo interface { //定义接口
	Name() string   //打印名字
	Date() string   //打印日期
	Gender() string //打印性别
}

func OnlyPrintNameAndGender(p PrintUserInfo) { //定义接口类型的变量作为参数
	fmt.Println(p.Name())
	fmt.Println(p.Gender())
	fmt.Println(p.Date())
}

func (s Student) PrintUserInfo() { //为结构体定义方法
	fmt.Println("name:", s.name)
	fmt.Println("date:", s.date)
	fmt.Println("gender:", s.gender)
}

type Foo int

func (f Foo) PrintVal() { //为内建类型定义方法
	fmt.Printf("this is %d\n", f)
}

func main() {
	s := Student{
		name:   "ahaoozhang",
		gender: "man",
		date:   time.Now(),
	}

	s.PrintUserInfo()

	var n Foo = 1
	n.PrintVal()

	OnlyPrintNameAndGender(&s)
}

// type S struct{ i int }

// func (p *S) Get() int  { return p.i }
// func (p *S) Put(v int) { p.i = v }

// type I interface {
// 	Get() int
// 	Put(int)
// }

// func f(p I) {
// 	fmt.Println(p.Get())
// 	p.Put(1)
// }

// func main() {
// 	var s S
// 	f(&s)
// }

package main

import . "fmt"
import "time"

type student struct {
	name string
	id   int32
	class
}

type class struct {
}

func (cl class) GetClass(a int32) string {
	return "hello class"
}

func (st student) GetName(a int32) string {
	return st.name
}

type Itfe interface {
	Func1()
	Func2()
	Func3()
}

func (cl class) Func1() string {
	return "Class Func1"
}
func (cl class) Func2() string {
	return "Class Func2"
}
func (cl class) Func3() string {
	return "Class Func3"
}

func Add(num *int32) {
	for i := 0; i < 1000; i++ {
		*num = *num + 1
	}
}

func main() {
	var st student
	var cl class
	st.id = 1
	st.name = "ahaoo"
	Println(st.GetName(5))
	Println(st.GetClass(5))
	Println(cl.GetClass(5))

	Println(cl.Func1())
	Println(cl.Func2())
	Println(cl.Func3())

	var num int32 = 0
	go Add(&num)
	go Add(&num)

	time.Sleep(time.Second)
	Println("num =", num)

}

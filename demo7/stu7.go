package main

import . "fmt"
import . "go_code/stu7/test"

func func1(str string) {
	Println(str)
}

func func2(str string) {
	Println("hhhhhhhhhhhhhhhhh")
}

func main() {

	funcptr := func1
	str := "ahaoozhang"
	funcptr(str)
	funcptr = func2
	funcptr(str)
	Println("hello world!")
	Test()

	var hf func(string)

	hf = func1
	hf(str)
}

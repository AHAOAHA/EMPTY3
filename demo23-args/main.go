package main

import "fmt"

//可变参数列表
func ArgsTestFunc(firstarg string, args ... string) {
	fmt.Println("firstarg:", firstarg)

	fmt.Printf("args: ")
	for _, v := range args {
		fmt.Printf("%s ", v)
	}
}

func main() {
	ArgsTestFunc("first", "a", "h", "a", "o", "o", "z", "h", "a", "n", "g")
}

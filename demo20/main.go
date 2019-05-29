package main

import "fmt"

type MyString struct {
	ptr *[]byte
	len int
}


func main() {
	s1 := "hello"
	s2 := "world"
	//s2 := s1
	//s2 = "hel"
	//fmt.Println(s1)
	//fmt.Println(s2)
	////s1[1] = 'd'	//不支持
	//for _, v := range s1 {
	//	fmt.Printf("%c\n", v)
	//}
	s3 := s1 + s2
	fmt.Println(s3)

}
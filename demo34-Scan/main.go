package main

import (
	"fmt"
	"bufio"
	"golang.org/x/net/html/atom"
	"os"
)

// fmt
func main() {
	/*
	var str string

	fmt.Scanln(&str)

	if len(str) == 0 {
		fmt.Println("len(str) == 0")
	}

	 */
	//var str string
	//var by []byte
	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Printf("input$ ")
	//str, _ = inputReader.ReadString('\n')
	//time.Sleep(time.Second*3)
	//n, _ := inputReader.Read(by)
	//fmt.Println("n:", n)
	//
	//fmt.Println(str)


	//var str []byte
	//
	//inputReader := bufio.NewReader(os.Stdin)
	//
	//
	//str,_ = inputReader.ReadBytes('\n')
	//
	//fmt.Println(str)
	//fmt.Printf("%+v", str)
	//fmt.Println("Buffered:", inputReader.Buffered())

	var str []byte
	inputReader := bufio.NewReader(os.Stdin)

	str, isPreFix, _ := inputReader.ReadLine()

	fmt.Println("str:", str)
	fmt.Println("isPreFix:", isPreFix)

}

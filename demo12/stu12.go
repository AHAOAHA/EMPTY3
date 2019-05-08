package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将结果放入chan C中
}

func RecvFromChan(c chan string) {
	time.Sleep(time.Second * 5)
	str := <-c //从chan中接收数据
	fmt.Println("Recv Success>", str)
}

func SendToChan(s string, c chan string) {
	//time.Sleep(time.Second * 5) //先休眠1s
	c <- s
	fmt.Println("Send Success>")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//chan
func main() {
	// s := []int{1, 2, 3, 4, 5, 6}
	// c := make(chan int) //定义可以用来传送int数据的chan C
	// go sum(s[len(s)/2:], c)
	// go sum(s[:len(s)/2], c)

	// x, y := <-c, <-c //x,y的结果为6， 15
	// //即可说明，chan可以被看作一个先进先出的管道
	// //fmt.Println(x, y, x+y)
	//-------------------------------------------------
	// str := "hello world!"
	// c := make(chan string) //定义用来传送string的chan
	// go RecvFromChan(c)
	// go SendToChan(str, c)
	// time.Sleep(time.Second * 10)
	//-------------------------------------------------

	//-------------------------------------------------
	// //select test
	// c := make(chan int)
	// quit := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}

	// 	quit <- 0
	// }()

	// fibonacci(c, quit)
	//-------------------------------------------------

	c := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		c <- 10
	}()
	for {
		select {
		case <-c:
			fmt.Println("x <- c is Ok!")
			panic("select")
		default:
			fmt.Println("default")
		}
	}

}

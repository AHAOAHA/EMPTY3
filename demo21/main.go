package main

import "fmt"

//map使用demo

func main() {
	//var mm map[string]string = {"zhango":"ahaoozhang"}	//未初始化的map不能使用
	mm := map[string]string {
		"zhango":"ahaoozhang",
	}
	//mm := make(map[string]string)
	
	mm [ "zhanghao" ] = "ahaoozhang"
	_, ok := mm["ahaoozg"]
	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	for k, v := range mm {
		fmt.Println(k, ":", v)
	}
}

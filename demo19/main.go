package main

//string结构

//type string struct{
//	len uint32
//	ptr *byte
//}

 type test struct{
	a int
}

func main() {
	////因为这里扩容是ptr在扩容，所以看不到&s发生变化
	//var s string = "hel"
	//fmt.Println(s, "&s:", &s, "len(s):", len(s))
	//s = "hello world"
	//fmt.Println(s, "&s:", &s, "len(s):", len(s))
	//var s1 string = "ahaoozhang"
	//s = s + s1
	//fmt.Println("s+s1:", s, "&s:", &s)	//+操作会先申请buf,拷贝原来的string，返回新的实例,这里仅设计实际存储字符串的buf的拷贝
	////t := test{}
	//
	//arr := [3]int32{1,2,3}
	//arr[1] = 0
	//for _, v := range arr {
	//	fmt.Printf("%d\n", v)
	//}

	sli := []int{}
	sli = append(sli, 1)

}

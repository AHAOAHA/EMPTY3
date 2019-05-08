package main

import (
	"fmt"
	"unsafe"
)

type sst1 struct {
	//定义结构体不需要var
	_a int
	_b string
}

func main() {
	//数组
	//数组声明
	//基本格式：var 数组名 [元素个数]元素类型

	//先声明再赋值
	var arr1 [5]int
	
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	//直接初始化
	var arr2 = [5]int{2, 4, 5, 1, 2}
	for k, v := range arr2 {
		fmt.Printf("arr2[%d]: %d\n", k, v)
	}

	//切片：[)前闭后开
	//切片指从一段连续内存中截取指定长度的元素
	//基本格式：[要进行切片的对象][开始位置:结束位置]
	//slice扩容机制类似于Linux下的vector机制
	var arr3 = [5]int{1, 2, 3, 4, 5}
	//普通切片
	fmt.Println(arr3[1:3])
	//开始位置缺省：表示从开始位置开始切片
	fmt.Println(arr3[:3])
	//结束位置缺省：表示结束位置即为数组边界
	fmt.Println(arr3[1:])
	//开始位置&结束位置同时缺省：表示数组自身
	fmt.Println(arr3[:])
	//两者同时为0：空切片,切出来的结果为空
	fmt.Println(arr3[0:0])
	//仅切最后一个位置时，可以使用len函数
	fmt.Println(arr3[len(arr3)-1])
	//切片范围越界：产生报错
	//fmt.Println(arr3[0:10])

	//切片声明
	//基本格式：var 切片名 []切片类型

	//声明含有多个未知元素的切片
	var sic []int
	fmt.Println(unsafe.Sizeof(sic)) //打印声明的切片大小为24
	fmt.Println(sic)
	//sic = arr3[1:3] 在对sic切片进行赋值操作之后，它的值不能再被认为是零值
	fmt.Println(sic == nil) //判定sic切片是否为空
	//sic切片未进行初始化，所以其值为编译器默认的零值

	//声明空切片
	var epsic = []int{}
	fmt.Println(unsafe.Sizeof(epsic)) //打印所声明的空切片的大小为24
	fmt.Println(epsic == nil)         //判定epsic切片是否为空
	//nil代表零值
	//epsic定义时初始化为空，则它的值不能被认为是零值

	//对nil值的测试
	//指针
	var ptr1 *int
	fmt.Println(ptr1 == nil) //--->true

	//结构体
	//var st1 sst1
	//fmt.Println(st1 == nil) 不支持该类型比较
	//截至目前 nil值仅支持指针、切片的比较

	//对切片增容的测试
	//扩容机制导致结果不相同
	s1 := []int{5}
	s1 = append(s1, 7)
	s1 = append(s1, 9)
	x1 := append(s1, 11)
	y1 := append(s1, 12)
	//y1
	fmt.Println(s1)
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println("---------------")
	s2 := []int{5, 7, 9}
	x2 := append(s2, 11)
	y2 := append(s2, 12)
	fmt.Println(s2)
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println("---------------")
	s3 := []int{7, 8, 9}
	fmt.Println(append(s3, 10))
	fmt.Printf("&s3: %p\n", &s3)
	fmt.Println(append(s3, 12))
	fmt.Println(append(s3, 12))
	fmt.Println(append(s3, 12))
	fmt.Println(append(s3, 12))
	fmt.Println(append(s3, 12))
	fmt.Printf("&s3: %p\n", &s3)
	fmt.Println(s3)

	//指针
	var a int
	fmt.Printf("&a:\n\%p: %p\n\%x:", &a)

}

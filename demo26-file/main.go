package main
//golang 文件操作

import (
	"fmt"
	"log"
	"os"
)

//创建文件
func CreateFile(filename string) {
	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//获取文件属性信息
func GetFileInfo(filename string) {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Size())
}


func main() {
	CreateFile("test.txt")
	GetFileInfo("text.txt")
}

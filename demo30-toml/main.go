package main
//使用golang 解析toml配置文件

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

//type songInfo struct {
//	Name string
//	Duration int
//}
//
//
//type config struct {
//	Bc string
//	Song songInfo
//}
//
//func test_toml() {
//	var cg config
//	var cpath string = "./config.toml"
//
//	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%v %v\n", cg.Bc, cg.Song.Name)
//}
//
//func main() {
//	test_toml()
//}


type MyConfig struct { // 整个文件可以看作一个大的结构体
	Title string // 全局字段可以直接提取
	User UserInfo // 每个[]对应一个结构体
	DataBase DataBaseInfo
	Servers map[string]Server
	Clients Clients
}

type UserInfo struct {
	Name string
	ID uint32
}

type DataBaseInfo struct {
	IP string
	Port uint64
	MaxConn uint32
}

type Server struct {
	IP string
	DC string
}

type Clients struct {
	Data [][]interface{}
	Hosts []string
}

func main() {
	var config MyConfig

	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		panic(err)
	}


	fmt.Printf("%v", config)
}
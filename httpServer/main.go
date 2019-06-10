package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

type myHandler struct {

}
func (this myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//对url进行解析
	u, err := url
	fmt.Fprintf(w, "hello world!\n")
}

//func myHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello world!\n")
//}

func main() {
	server := http.Server {
		Addr: ":8080",
		Handler: &myHandler{},
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe)
}

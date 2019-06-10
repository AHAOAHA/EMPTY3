package main

import (
	"github.com/json-iterator/go"
	"os"
)

type Color struct {
	ID int
	Name string
	Color []string
}

func main() {
	data := Color{
		ID: 1,
		Name: "red",
		Color: []string{"ahaoozhang", "ayuan"},
	}

	b, _ := jsoniter.Marshal(data)

	os.Stdout.Write(b)
}

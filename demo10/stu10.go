package main

import "fmt"

type student struct {
	id int32
	name string
}

func main() {
	ahaoo := &student{
		id:	1,
		name:	"ahaoo",
	}
	ahao := student{
		id: 2,
		name:	"ahao",
	}

	ajing := student {
		3,
		"ajing",
	}

	aqiang := &struct {
		id int32
		name string
	}{
		4,
		"aqiang",
	}

	fmt.Println(ahaoo.id, ahaoo.name)
	fmt.Println(ahaoo.id, ahaoo.name)
	fmt.Println(ahaoo)
	fmt.Println(ahao)
	fmt.Println(ajing)
	fmt.Println(aqiang)

}

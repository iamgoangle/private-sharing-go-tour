package main

import (
	"fmt"
)

type Employee struct {
	firstname, lastname string
}

var m map[string]Employee

func main() {
	m = make(map[string]Employee)
	m["developer"] = Employee{
		"Teerapong", "Singthong",
	}

	fmt.Println(m)
}

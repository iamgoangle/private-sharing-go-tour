package main

import (
	"fmt"
)

type Employee struct {
	firstname string
	lastname  string
}

func (e Employee) GetFullname() string {
	return e.firstname + "\t" + e.lastname
}

func main() {
	var newE = Employee{
		firstname: "Teerapong",
		lastname:  "Singthong",
	}

	fmt.Println(newE.GetFullname())
}

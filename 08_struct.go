package main

import (
	"fmt"
)

type Employee struct {
	firstname string
	lastname  string
	age       int
	hire      bool
}

func main() {
	var newEmp = Employee{
		firstname: "Teerapong",
		lastname:  "Singthong",
		age:       28,
		hire:      true,
	}

	fmt.Println("New Employee")
	fmt.Printf("Firstname = %v \n", newEmp.firstname)
	fmt.Printf("Lastname = %v \n", newEmp.lastname)
	fmt.Printf("Age = %v \n", newEmp.age)
	fmt.Printf("Hire = %v \n", newEmp.hire)
}

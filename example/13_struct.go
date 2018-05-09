package main

import "fmt"

type Employee struct {
	firstname string
	lastname  string
}

func main() {
	var newEmp = Employee{
		firstname: "Teerapong",
		lastname:  "Singthong",
	}

	// or new keyword
	var newProgrammer = new(Employee)
	newProgrammer.firstname = "David"
	newProgrammer.lastname = "Kopper Field"

	fmt.Printf("Firstname = %v \n", newEmp.firstname)
	fmt.Printf("Lastname = %v \n", newEmp.lastname)

	fmt.Printf("Firstname = %v \n", newProgrammer.firstname)
	fmt.Printf("Lastname = %v \n", newProgrammer.lastname)
}

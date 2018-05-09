package main

import (
	"fmt"
)

type Employee struct {
	firstname, lastname string
}

var m = map[string]Employee{
	"developer": Employee{
		firstname: "Teerapong",
		lastname:  "Singthong",
	},
}

func main() {
	fmt.Println(m)
}

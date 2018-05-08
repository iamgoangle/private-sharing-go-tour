package main

import (
	"fmt"
)

func main() {
	const company = "Line corp"
	var i, j bool = true, false
	shortHand := "Teerapong Singthong"

	fmt.Printf("Type = %T Value = %v \n", i, i)
	fmt.Printf("Type = %T Value = %v \n", j, j)
	fmt.Printf("Type = %T value = %v \n", shortHand, shortHand)
	fmt.Printf("Type = %T Value = %v", company, company)
}

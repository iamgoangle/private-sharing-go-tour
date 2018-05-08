package main

import (
	"fmt"
)

func avg(num []int) int {
	return num[0]
}

func main() {
	x := []int{
		1,
		2,
		3,
		4,
		5
	}

	fmt.Print(avg(x))
}

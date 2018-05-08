package main

import "fmt"

func main() {
	var list [5]int
	list[4] = 100
	fmt.Println(list)

	x := [4]int{
		1,
		2,
		3,
		4,
	}

	fmt.Println(x)
}

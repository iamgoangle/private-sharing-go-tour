package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer counting stack")

	for i := 0; i < 1000; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done...")
}

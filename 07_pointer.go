package main

import "fmt"

func zero(xPrt *int) {
	*xPrt = 0 // refs to pointer
}

func main() {
	fmt.Println("Before pointer")
	x := 5
	fmt.Println(x)

	zero(&x) // generate pointer
	fmt.Println("After pointer")
	fmt.Println(x)
}

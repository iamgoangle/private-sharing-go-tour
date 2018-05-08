package main

import "fmt"

func forLoop() {
	fmt.Println("forLoop")

	var i int
	for i = 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	fmt.Println("whileLoop")

	var i int = 1
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
}

func main() {
	forLoop()
	whileLoop()
}

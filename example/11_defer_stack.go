package main

import (
	"fmt"
)

func hello() {
	fmt.Print("Hello 1")
}

func world() {
	fmt.Println("World 2")
}

func main() {
	defer world()
	hello()
}

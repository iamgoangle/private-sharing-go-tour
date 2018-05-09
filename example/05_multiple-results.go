package main

import (
	"fmt"
)

func swap(x, y string) (string, int) {
	return x + y + "   ", 1234
}

func main() {
	fmt.Print(swap("g", "o"))
}

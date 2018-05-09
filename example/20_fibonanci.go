package main

import (
	"fmt"
)

func fibo(n int) int {
	for n > 0 {

		if n < 2 {
			return 1
		} else {
			return fibo(n-2) + fibo(n-1)
		}

		n--
	}

	return 0
}

func main() {
	fmt.Println(fibo(6))
}

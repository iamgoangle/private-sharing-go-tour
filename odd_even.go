package main

import "fmt"

func oddEven() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Println(" even")
		} else {
			fmt.Println(" odd")
		}
	}
}

func main() {
	oddEven()
}

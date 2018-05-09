package main

import (
	"fmt"
)

func mutable(i *int) {
	*i = 10 // change value in ddress
}

func main() {
	i := 40

	mutable(&i) // ref i address

	fmt.Println(i) // 10
}

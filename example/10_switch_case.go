package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Print("OS X")
	case "linux":
		fmt.Print("Linux")
	default:
		fmt.Print("%s", os)
	}
}

package main

import "fmt"

func main() {
	x := make(map[string]string)

	x["test"] = "hello world"
	x["firstname"] = "teerapong"
	x["lastname"] = "singthong"

	fmt.Println(x)

	delete(x, "test") // delete key test

	fmt.Println(x)

	// defaukt map value
	group := map[string]string{
		"job":    "Software Engineer",
		"salary": "100,000",
	}

	fmt.Println(group)
}

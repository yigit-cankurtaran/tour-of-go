package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	// this is a type assertion
	fmt.Println(s)

	s, ok := i.(string)
	// ok is true if the assertion holds
	// and false otherwise
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	// because i is not a float64
	fmt.Println(f)
}

package main

import "fmt"

func main() {
	var i interface{}
	// i is an empty interface
	// these are used by code that handles values of unknown type
	describe(i)

	i = 42
	describe(i)
	// prints: (42, int)

	i = "hello"
	describe(i)
	// prints: (hello, string)
	// the interface can hold any type, is dynamic
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

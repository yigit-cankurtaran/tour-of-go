package main

import "fmt"

// this is a function that takes two ints and returns an int
// this is the syntax for declaring a function
func add(x, y int) int {
	// we can shorten x int, y int to x, y int
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

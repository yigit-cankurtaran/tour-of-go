package main

import "fmt"

// this is a function that takes two ints and returns an int
// this is the syntax for declaring a function
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

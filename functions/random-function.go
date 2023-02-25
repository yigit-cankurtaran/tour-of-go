package main

import (
	"fmt"
	"math/rand"
)

// this is a function that takes two ints and returns an int
// this is the syntax for declaring a function
func add(x int, y int) int {
	return x + y
}

// adding 2 random numbers together
func main() {
	fmt.Println(add(rand.Int(), rand.Int()))
}

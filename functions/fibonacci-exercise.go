package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.

func fibonacci() func() int {
	a, b := 0, 1
	// a and b are the two numbers in the fibonacci sequence
	return func() int {
		a, b = b, a+b
		// this is the fibonacci sequence
		// a is the current number
		// b is the next number
		// the next number is the sum of the current number and the previous number
		return a
		// returns the current number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

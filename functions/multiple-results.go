package main

import "fmt"

func swap(x, y string) (string, string) {
	// takes 2 strings, returns 2 strings
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	// := is a short assignment statement
	// foo := 42 is the same as var foo int = 42
	fmt.Println(a, b)
}

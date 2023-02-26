package main

import "fmt"

func main() {
	var a [2]string
	// a is an array of 2 strings
	// the type of a is [2]string
	// length is part of the type
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	// prints Hello World
	fmt.Println(a)
	// prints [Hello World]
	// because a is an array, not a slice
	a[0] = "Goodbye"
	fmt.Println(a[0], a[1])
	// prints Goodbye World
	fmt.Println(a)
	// prints [Goodbye World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// prints [2 3 5 7 11 13]
}

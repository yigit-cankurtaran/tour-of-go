package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// primes is an array of 6 ints, fixed length

	var s []int = primes[1:4]
	// s is a slice of primes
	// s has a length of 3
	// prints elements 1, 2, 3

	fmt.Println(s)
	// prints [3 5 7]
}

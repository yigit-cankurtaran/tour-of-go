package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// extend its length
	s = s[:4]
	printSlice(s)
	// this changes the underlying array, so the next slice will be different

	// drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// what this does is print the length, capacity, and all of the elements
}

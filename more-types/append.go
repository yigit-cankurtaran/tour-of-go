package main

import "fmt"

func main() {
	var s []int
	printSlice(s)
	// nil slice

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)
	// prints: len=1 cap=1 [0]

	s = append(s, 1)
	printSlice(s)
	// prints: len=2 cap=2 [0 1]

	// we can add multiple elements with one call
	s = append(s, 2, 3, 4)
	// how this syntax works is that the first argument is the slice
	// and the rest are the elements to append
	printSlice(s)
	// the capacity is doubled when the slice is full
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

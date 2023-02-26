package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	// q is a slice of 6 ints, variable length
	// [] means slice, not array
	// a slice literal is like an array literal without the length
	// creates the same array, builds a slice that references it
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	// r is a slice of 6 bools
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	// s is a slice of 6 structs, each with an int and a bool
	fmt.Println(s)
	// the struct values are printed in the same order as the struct fields
}

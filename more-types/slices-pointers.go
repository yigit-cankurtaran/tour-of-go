package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
		// the beatles lol
	}
	fmt.Println(names)

	a := names[0:2]
	// prints elements 0, 1
	b := names[1:3]
	// prints elements 1, 2
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	// prints [John XXX] [XXX George]
	fmt.Println(names)
	// prints [John XXX George Ringo]
	// changed the entire array, because slices don't have their own data
	// they are just pointers to the underlying array
}

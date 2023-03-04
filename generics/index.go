package main

import "fmt"

// Index returns the first index of the target string t, or -1 if no match is found.
func Index[T comparable](s []T, x T) int {
	// this function is generic, it can be used with any type that is comparable
	// declaration means s is a slice of type T, and x is of type T
	// as long as T is comparable, the function can be used with any type that is comparable
	// comparable means that the type can be compared with ==, !=, <, <=, >, >=
	for i, v := range s {
		if v == x {
			// v and x have the same type, so they can be compared
			return i
			// return the index of the first match
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index works on a slice of strings
	ss := []string{"a", "b", "c", "d"}
	fmt.Println(Index(ss, "c"))
}

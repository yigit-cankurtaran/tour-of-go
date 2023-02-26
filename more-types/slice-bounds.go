package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)
	// prints elements 1, 2
	// default is 0 for the low bound

	s = s[1:]
	fmt.Println(s)
	// prints element 2
	// default is len(s) for the high bound
}

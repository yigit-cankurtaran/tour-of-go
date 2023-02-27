package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
		// this line prints the index and the value of the element
	}
	// this loops until the end of the slice
	// range returns both the index and the value of the element
}

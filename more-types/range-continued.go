package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		// this is the same as pow[i] = 2**i
	}
	for _, value := range pow {
		// we can use the blank identifier to ignore the index
		fmt.Printf("%d\n", value)
	}
}

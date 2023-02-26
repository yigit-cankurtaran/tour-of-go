package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		// what this does is it stacks the defers
		// the earliest defer is 9, the last is 0
		// so, it prints 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
	}

	fmt.Println("done")
	// starts printing after this
}

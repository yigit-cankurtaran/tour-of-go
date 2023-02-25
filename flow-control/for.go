package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		// i is 0, while i is less than 10, add 1 to i, and add i to sum
	}
	fmt.Println(sum)
}

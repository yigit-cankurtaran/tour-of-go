package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		// the init and post statements are omitted
		// the code still works properly
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
	// the earliest defer is the last to be executed
	// defers the execution of a function until the surrounding function returns

	fmt.Println("hello")
}

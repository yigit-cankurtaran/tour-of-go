package main

import "fmt"

const Pi = 3.14

// constants can be character, string, boolean, or numeric values
// constants CAN NOT be declared using the := syntax

func main() {
	const World = "世界"
	// the chinese characters above mean "world"
	// each are 3 bytes
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

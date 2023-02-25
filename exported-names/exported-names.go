package main

// capital letters denote exported names

import (
	"fmt"
	"math"
)

// Pi is exported, pi is not
// this prints the first few digits of pi
func main() {
	fmt.Println(math.Pi)
}

// func main() {
// fmt.Println(math.pi)
// }
// this prints an error because pi is not exported

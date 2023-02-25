package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// the code above squares x and y, adds them together, and then takes the square root of the sum
	// then saves the sum as a float64 in f
	var z uint = uint(f)
	// the code above converts f to a uint and saves it in z
	fmt.Println(x, y, z)
}

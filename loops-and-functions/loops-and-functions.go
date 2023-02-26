package main

import "fmt"

/*
create a program that can calculate the square root of a number
using a for loop. The loop starts with a guess of 1.0, and then
improves the guess by averaging it with the quotient of x/guess.
The loop stops when the i hits the number in the for loop.
*/

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		// Newton's method
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1100))
	// tried with lots of numbers, 10 seems to be either correct or close enough
}

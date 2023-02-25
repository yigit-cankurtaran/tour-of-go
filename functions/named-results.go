package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	// this function returns 2 ints
	// the return statement without arguments returns the named return values
	// the "named" returns are the x and y in the function declaration
	// this is known as a "naked" return
	// naked returns should be used only in short functions
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split(63))
}

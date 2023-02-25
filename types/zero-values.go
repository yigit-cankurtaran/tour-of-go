package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// string uses %q to print empty string
	// the q stands for quoted
	// the others use %v to print zero value
	// the v stands for value
}

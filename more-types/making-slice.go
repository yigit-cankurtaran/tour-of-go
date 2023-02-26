package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	// using this syntax, you can specify the length and capacity
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	// the capacities by default are maxed out at the length of the underlying array
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

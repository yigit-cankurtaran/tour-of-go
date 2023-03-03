package main

import "fmt"

type I interface {
	M()
	// this is a nil interface value
	// it holds neither value nor concrete type
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	// gives a runtime error: invalid memory address or nil pointer dereference
	// because i is nil, there is no value or concrete type to print
}

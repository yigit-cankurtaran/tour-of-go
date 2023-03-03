package main

import (
	"fmt"
	"math"
)

type I interface {
	// interfaces are a tuple of a value and a concrete type
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// this is a method with a pointer receiver
	// implements I
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	// this is a method with a value receiver
	// implements I
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

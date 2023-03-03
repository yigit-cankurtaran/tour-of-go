package main

import "fmt"

// stringer is an interface in the fmt package
// it is implemented by any value that has a String method
// String method returns a string

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	// prints Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}

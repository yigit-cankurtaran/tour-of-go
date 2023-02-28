package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// insert or update an element in map m

	delete(m, "Answer")
	// delete removes the key/value pair from the map
	fmt.Println("The value:", m["Answer"])
	// prints 0, the default value for int

	v, ok := m["Answer"]
	// tests if a key is present in a map
	// if key is in m, ok is true. if not, ok is false
	fmt.Println("The value:", v, "Present?", ok)
	// prints 0 false
}

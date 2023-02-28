package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// a map maps keys to values.
// it can be used to implement associative arrays or dictionaries or hashes.

func main() {
	m = make(map[string]Vertex)
	// make returns a map of the given type
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

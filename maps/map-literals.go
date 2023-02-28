package main

// map literals are like struct literals
// BUT we need the keys

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
	// we could remove the vertex statement and just have the values
}

func main() {
	fmt.Println(m)
}

// the map is a reference to the data structure created by make

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	// prints the bounds of the image and the RGBA values of the pixel at (0, 0).
	// the bounds are (0, 0) to (100, 100), and the RGBA values are (0, 0, 0, 0).
}

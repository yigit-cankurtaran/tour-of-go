package main

import (
	"image"
	"image/color"
	// "golang.org/x/tour/pic"
	// these imports don't work locally, i'll run it on the playground
)

type Image struct {
	width, height int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// m := Image{256, 256}
	// pic.ShowImage(m)
	// returned a picture, works perfectly
}

package main

// import "golang.org/x/tour/pic"

// the import didn't work locally, will try running it on the website
// the website gives me a picture. this is a success.

func Pic(dx, dy int) [][]uint8 {
	// create the outer slice
	outer := make([][]uint8, dy)
	// create the inner slices
	for i := range outer {
		outer[i] = make([]uint8, dx)
	}
	// initialize the inner slices
	for i := range outer {
		for j := range outer[i] {
			outer[i][j] = uint8(i * j)
		}
	}
	return outer
}

func main() {
	// pic.Show(Pic)
}

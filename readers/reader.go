package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	// creates a new Reader reading from a string
	b := make([]byte, 8)
	// creates a byte slice with a length of 8
	for {
		n, err := r.Read(b)
		// reads the next len(b) bytes from the Reader, storing them in b.
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// prints the number of bytes read, the error if any, and the bytes as a slice of bytes.
		fmt.Printf("b[:n] = %q\n", b[:n])
		// prints the bytes read as a string.
		if err == io.EOF {
			break
			// if the error is io.EOF, the end of the Reader has been reached.
		}
	}
}

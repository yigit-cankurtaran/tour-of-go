package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	// using a pointer receiver because the Read method needs to modify the slice
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		// for each byte in the slice
		c := b[i]
		// c is the current byte
		if (c >= 'A' && c <= 'M') || (c >= 'a' && c <= 'm') {
			b[i] += 13
			// if the byte is between A-M or a-m, add 13 to it
		} else if (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z') {
			b[i] -= 13
			// if the byte is between N-Z or n-z, subtract 13 from it
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// prints "You cracked the code!"
	r := &rot13Reader{s}
	io.Copy(os.Stdout, r)
}

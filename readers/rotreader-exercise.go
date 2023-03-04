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
		c := b[i]
		if (c >= 'A' && c <= 'M') || (c >= 'a' && c <= 'm') {
			b[i] += 13
		} else if (c >= 'N' && c <= 'Z') || (c >= 'n' && c <= 'z') {
			b[i] -= 13
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

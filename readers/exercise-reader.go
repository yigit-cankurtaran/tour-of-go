package main

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// MyReader implements the Reader interface
	for i := range b {
		// b is a slice of bytes, for each byte in b
		b[i] = 'A'
		// sets each byte to 'A'
	}
	return len(b), nil
	// the import doesn't work locally, i'll run it on the playground
}

func main() {
	reader.Validate(MyReader{})
	// got the "OK!" message, works as intended.
}

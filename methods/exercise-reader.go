package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(p []byte) (int, error) {
	p[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

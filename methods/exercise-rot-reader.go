package main

import (
	"io"
	"os"
	"strings"
)

func rot13(r byte) byte {
	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r < 'm' {
			return r + 13
		} else {
			return r - 13
		}
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r < 'M' {
			return r + 13
		} else {
			return r - 13
		}
	}
	// Do nothing.
	return r
}

type rot13Reader struct {
	r io.Reader
}

func (a rot13Reader) Read(p []byte) (int, error) {
	n, err := a.r.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = rot13(p[i])
	}
	copy(p, buf)
	return len(p), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

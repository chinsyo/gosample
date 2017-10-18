package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, e := rot.r.Read(b)

	for i := 0; i < len(b); i++ {
		switch {
		case 'a' <= b[i] && b[i] <= 'm':
			fallthrough
		case 'A' <= b[i] && b[i] <= 'M':
			b[i] = b[i] + 13
		case 'n' <= b[i] && b[i] <= 'z':
			fallthrough
		case 'N' <= b[i] && b[i] <= 'Z':
			b[i] = b[i] - 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

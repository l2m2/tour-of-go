package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (v *rot13Reader) Read(b []byte) (int, error) {
	n, err := v.r.Read(b)
	for i, val := range b {
		if val >= 'A' && val <= 'Z' {
			val = (val-'A'+13)%26 + 'A'
		} else if val >= 'a' && val <= 'z' {
			val = (val-'a'+13)%26 + 'a'
		}
		b[i] = val
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

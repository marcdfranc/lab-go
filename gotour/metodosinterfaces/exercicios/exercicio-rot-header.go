package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, nil
}

func rot13(b byte) byte {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
			b += 13
		} else {
			b -= 13
		}
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

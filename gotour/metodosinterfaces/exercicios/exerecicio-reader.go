package main

import (
	"fmt"
	"io"
	"os"
)

type infiniteAReader struct{}

// Read preenche o slice de bytes passado com o caractere 'A'
func (r infiniteAReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader := infiniteAReader{}
	buffer := make([]byte, 8) // Buffer de exemplo com 8 bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Erro:", err)
			os.Exit(1)
		}
		fmt.Print(string(buffer[:n]))
	}
}

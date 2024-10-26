// remove eleemntos de navegação em diretorios e sufixos de extensão de arquivo
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s\n", basename(input))
}

func basename(s string) string {

	// remove tudo que estiver antes da ultima "/" encontrada
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

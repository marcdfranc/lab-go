// remove eleemntos de navegação em diretorios e sufixos de extensão de arquivo
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s\n", basename2(input))
}

func basename2(s string) string {

	slash := strings.LastIndex(s, "/") // se nao encontrar retorna -1
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

// Exibe a contagem e o texto das que aparecem mais de uma
// vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		fmt.Println(files)
		for _, arg := range files {
			f, err := os.Open(arg)

			fName := strings.Split(arg, "\\")

			if err != nil {
				fmt.Printf("dup2: %v\n", err)
				continue
			}
			countLines(f, fName[len(fName)-1], counts)
			f.Close()
		}
	}

	for item, n := range counts {
		if n > 1 {
			line := strings.Split(item, "|")

			fmt.Printf("File %s: %s\t%d\n", line[0], line[1], n)
		}
	}
}

func countLines(f *os.File, fName string, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[fName+"|"+input.Text()]++
	}
}

// Exibe o texto de toda linha que aprece mais de uma vez na entrada padrao,
// preceidida por sua contagem\
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTA: igonorando erros potencias com input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\\%s\n", n, line)
		}
	}
}

package main

import (
	"fmt"
)

// Função para contar os bits 1
func CountBits(num uint64) int {
	count := 0
	for num != 0 {
		count += int(num & 1) // Incrementa o contador se o bit menos significativo for 1
		num >>= 1             // Desloca os bits para a direita
	}
	return count
}

func main() {
	num := uint64(299536) // Exemplo de número: 29 (em binário: 11101)
	fmt.Printf("O número %d tem %d bits 1.\n", num, CountBits(num))
}

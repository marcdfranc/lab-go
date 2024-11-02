package main

import (
	"fmt"
)

func countSpaces(b []byte) int {
	count := 0
	for _, v := range b {
		if v == ' ' {
			count++
		}
	}
	return count
}

func main() {
	s := []byte("Olá, mundo! Como você está?")
	fmt.Printf("Número de espaços: %d\n", countSpaces(s))
}

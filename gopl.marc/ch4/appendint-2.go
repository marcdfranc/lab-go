// Append Int
package main

import "fmt"

func main() {
	var x []int

	for i := 0; i < 10; i++ {
		x = appendInt2(x, i+1, i+2, i+3, i+4, i+5)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(x), x)
	}
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)

	if zLen <= cap(x) {
		// Ha espaco para crescer
		z = make([]int, zLen, cap(x))
		copy(z[:len(x)], x)
		copy(z[len(x):], y)
	} else {
		// nao ha espaco psuficiente, aloca novo array
		// cresce para o dobro, para a complexidade linear amortizada

		zCap := zLen

		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z[:len(x)], x)
		copy(z[len(x):], y)
	}
	return z
}

// Append Int
package main

import "fmt"

func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1

	if zLen <= cap(x) {
		// Ha espaco para crescer
		z = x[:zLen]
	} else {
		// nao ha espaco psuficiente, aloca novo array
		// cresce para o dobro, para a complexidade linear amortizada

		zCap := zLen

		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

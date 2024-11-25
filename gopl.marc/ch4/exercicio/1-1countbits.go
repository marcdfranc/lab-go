package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%v\n%v\n similar bytes\n", c1, c2)
	fmt.Printf("Total: %d Bytes\n", countBytes(&c1, &c2))
}

func countBytes(a *[32]byte, b *[32]byte) int {
	var counter int
	for i := range a {
		if a[i] == b[i] {
			fmt.Println(a[i])
			counter++
		}
	}
	return counter
}

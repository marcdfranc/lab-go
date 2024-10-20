package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc)
	// fmt.Println(len(pc))
}

func main() {
	num := uint64(299536) // Exemplo de número: 29 (em binário: 11101)
	fmt.Printf("O número %d tem %d bits 1.\n", num, PopCount(num))
}

func PopCount(x uint64) int {

	var r int = 0
	fmt.Printf("%b\n", x)
	for i := 0; i < 8; i++ {
		fmt.Printf("%v\n", x>>(i*8))
		r += int(pc[byte(x>>(i*8))])
	}
	return r
	/*return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))])*/
}

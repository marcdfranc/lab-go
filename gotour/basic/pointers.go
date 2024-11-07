package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // ponteiro para i
	fmt.Println(*p) // Lê i através do ponteiro
	*p = 21         // atribui a i atraves do ponteiro
	fmt.Println(i)  // siplay o valor de i

	p = &j         // ponteiro para j
	*p = *p / 37   // divide j atraves do ponteiro
	fmt.Println(j) // display o novo valor de j
}

package main

import "fmt"

func main() {
	var a = [...]int{10, 20, 30}

	fmt.Println(a[0])        // Primeiro elemento
	fmt.Println(a[len(a)-1]) // ultimom elemento

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}

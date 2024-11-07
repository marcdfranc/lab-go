package main

import "fmt"

func main() {
	var s []int
	printSlice3(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice3(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice3(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5, 6)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println(WordCount("A man a plan a canal panama."))
}

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	fields := strings.Fields(s)

	for _, field := range fields {
		words[field]++
	}
	return words
}

package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "b", "c", "c", "c", "d", "d", "e"}
	removeAdjacentDuplicates(&s)
	fmt.Println(s)
}

func removeAdjacentDuplicates(strings *[]string) {
	if len(*strings) == 0 {
		return
	}

	j := 0
	for i := 1; i < len(*strings); i++ {
		if (*strings)[i] != (*strings)[j] {
			j++
			(*strings)[j] = (*strings)[i]
		}
	}

	*strings = (*strings)[:j+1]
}

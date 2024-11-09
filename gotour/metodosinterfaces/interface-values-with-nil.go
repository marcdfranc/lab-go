package main

import "fmt"

type I3 interface {
	M()
}

type T3 struct {
	S string
}

func (t *T3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I3

	var t *T3
	i = t
	describe2(i)
	i.M()

	i = &T3{"hello"}
	describe2(i)
	i.M()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

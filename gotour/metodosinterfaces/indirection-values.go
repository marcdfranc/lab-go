package main

import (
	"fmt"
	"math"
)

type Vertex9 struct {
	X, Y float64
}

func (v Vertex9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex9) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex9{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex9{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

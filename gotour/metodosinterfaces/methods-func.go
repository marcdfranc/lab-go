package main

import (
	"fmt"
	"math"
)

type Vertex6 struct {
	X, Y float64
}

func Abs(v Vertex6) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex6{3, 4}
	fmt.Println(Abs(v))
}

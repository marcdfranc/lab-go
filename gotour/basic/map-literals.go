package main

import "fmt"

type Vertex6 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex6{
	"Bell Labs": Vertex6{
		40.68433, -74.39967,
	},
	"Google": Vertex6{
		37.42202, -122.08408,
	},
}

var m3 = map[string]Vertex6{
	"Novo Lab": {40.68433, -74.39967},
	"Bing":     {37.42202, -122.08408},
}

func main() {
	fmt.Println(m2)
	fmt.Println(m3)
}

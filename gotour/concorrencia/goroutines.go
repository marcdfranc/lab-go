package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// esses caras executam em paralelo.
	go say("world")
	go say("hello")
	say("Marcelo")
}

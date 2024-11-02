package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	BRL
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", BRL: "R$"}
	fmt.Println(USD, symbol[USD])
	fmt.Println(EUR, symbol[EUR])
	fmt.Println(GBP, symbol[GBP])
	fmt.Println(BRL, symbol[BRL])
}

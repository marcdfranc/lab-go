package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	defaultTxt := "Default"

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "384":
			showSha384(defaultTxt)
		case "512":
			showSha512(defaultTxt)
		default:
			showSha256(defaultTxt)
		}
	} else {
		showSha256(defaultTxt)
	}
}

func showSha256(txt string) {
	v := sha256.Sum256([]byte(txt))
	fmt.Printf("%x\n", v)
}

func showSha384(txt string) {
	v := sha512.Sum384([]byte(txt))
	fmt.Printf("%x\n", v)
}

func showSha512(txt string) {
	v := sha512.Sum512([]byte(txt))
	fmt.Printf("%x\n", v)
}

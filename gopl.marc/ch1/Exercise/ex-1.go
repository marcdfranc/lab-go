// exibe comando e argumentos
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))

	fmt.Println(os.Args)
}

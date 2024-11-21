// main.go
package main

import (
	"cacheservice/logging"
	"os"
)

func main() {
	logger := logging.NewLogger(os.Stdout, "main.go")

	num := 2.5
	integer := 77

	type someObj struct {
		firstName string
		age       int
	}
	obj := someObj{"Maria", 38}

	logger.LogInfof("decimal: %f, int: %d, obj: %v", num, integer, obj)
}

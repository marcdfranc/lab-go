// main.go
package main

import (
	"cacheservice/httpserv"
	"cacheservice/logging"
	"fmt"
	"net/http"
	"os"
)

func main() {
	logger := logging.NewLogger(os.Stdout)

	num := 2.5
	integer := 77

	type someObj struct {
		firstName string
		age       int
	}

	obj := someObj{"Maria", 38}

	logger.LogInfof("decimal: %f, int: %d, obj: %v", num, integer, obj)

	serv := httpserv.NewHttpserv("localhost:8080", logger)

	serv.HandleFunc("/", idx)
	serv.HandleFunc("/getsome", getParam)

	serv.Start()
}

func idx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index page")
}

func getParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "request: %v", r)
}

// main.go
package main

import (
	"cacheservice/logger"
	"log"
	"net/http"
)

func main() {
	// Injeção de dependência do ConsoleOutput
	output := &logger.ConsoleOutput{}
	logInstance := logger.NewLogger(output, "main")

	// Uso do Logger na aplicação
	logInstance.LogInfo("Iniciando servidor...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logInstance.LogInfo("Recebida requisição para /")
		w.Write([]byte("Hello, World!"))
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

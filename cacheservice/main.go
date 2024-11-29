// main.go
package main

import (
	"cacheservice/handlers"
	"cacheservice/httpserv"
	"cacheservice/logging"
	"os"
)

func main() {
	logger := logging.NewLogger(os.Stdout)

	server := httpserv.NewHttpServer("localhost:8000", logger)
	handler := handlers.NewKeyValueHandler(logger, server)

	server.HandleGet("/api/keyvalue", handler.GetHandler)
	server.HandlePost("/api/keyvalue", handler.PostHandler)

	server.HandleGet("/api/keyvalue/{id}", handler.GetWithParamHandler)
	server.HandlePut("/api/keyvalue/{id}", handler.PutHandler)
	server.HandleDelete("/api/keyvalue/{id}", handler.DeleteHandler)

	server.Start()
}

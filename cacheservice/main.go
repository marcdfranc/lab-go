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

	server.HandleGet("/get", handler.GetHandler)
	server.HandleGet("/getp/{id}", handler.GetWithParamHandler)
	server.HandleHead("/head", handler.HeadHandler)
	server.HandlePost("/post", handler.PostHandler)
	server.HandlePut("/put", handler.PutHandler)
	server.HandlePatch("/patch", handler.PatchHandler)
	server.HandleDelete("/delete", handler.DeleteHandler)
	server.HandleConnect("/connect", handler.ConnectHandler)
	server.HandleOptions("/options", handler.OptionsHandler)
	server.HandleTrace("/trace", handler.TraceHandler)

	server.Start()
}

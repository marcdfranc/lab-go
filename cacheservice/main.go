package main

import (
	"cacheservice/api"
)

// var _logger = &logger.Logger{Name: "main.go"}

func main() {
	/*_logger.Log("My First event")
	_logger.LogInfo("My First info event")
	_logger.LogDebug("My First Debug event")
	_logger.LogWarning("My First Warn event")
	_logger.LogError("My First Error event", fmt.Errorf("could not process"))
	_logger.LogError("My First Error event", fmt.Errorf("Null Pointer"))
	_logger.LogError("My First Error event", fmt.Errorf("I don't know"))
	_logger.LogFatal("My First Fatal event", fmt.Errorf("Critical failure"))*/

	api.StartServer()
}

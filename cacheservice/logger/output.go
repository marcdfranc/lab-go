package logger

import (
	"fmt"
	"os"
	"time"
)

type IOutput interface {
	Log(level LogLevel, name, message string, errors ...error)
}

type ConsoleOutput struct{}

func (co *ConsoleOutput) Log(level LogLevel, name, message string, errors ...error) {
	if level == UNDEFINED {
		fmt.Println(message)
		return
	}
	fmt.Printf("[%s] Level: %s Time: %v Msg: %s\n", name, level, time.Now(), message)
	for i, err := range errors {
		fmt.Printf("Error (%d): %v\n", i, err)
	}
	if errors != nil && level == FATAL {
		os.Exit(1)
	}
}

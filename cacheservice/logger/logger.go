package logger

import (
	"fmt"
	"os"
	"time"
)

func (l *Logger) log(level LogLevel, message string, errors ...error) {
	if level == UNDEFINED {
		fmt.Println(message)
		return
	}
	fmt.Printf("[%s] Level: %s	Time: %v	Msg: %s\n", l.Name, level, time.Now(), message)
	for i, err := range errors {
		fmt.Printf("Error (%d): %v\n", i, err)
	}
	if errors != nil && level == FATAL {
		os.Exit(1)
	}
}

package logging

import (
	"fmt"
	"io"
	"time"
)

type Logging interface {
	LogInfo(message string)
	LogDebug(message string)
	LogWarning(message string)
	LogError(message string)
	LogInfof(formatter string, data ...interface{})
	LogDebugf(formatter string, data ...interface{})
	LogWarningf(formatter string, data ...interface{})
	LogErrorf(formatter string, data ...interface{})
}

type Logger struct {
	output io.Writer
}

func NewLogger(output io.Writer) *Logger {
	return &Logger{
		output: output,
	}
}

func (l *Logger) log(level LogLevel, template string, data ...any) {
	template = fmt.Sprintf("[%s] Time: %s, %s \n", level, time.Now().UTC().Format(time.RFC3339), template)
	_, err := fmt.Fprintf(l.output, template, data...)
	if err != nil {
		panic(err)
	}
}

func (l *Logger) LogInfo(message string) {
	l.log(INFO, "Msg: %s ", message)
}

func (l *Logger) LogDebug(message string) {
	l.log(DEBUG, "Msg: %s ", message)
}

func (l *Logger) LogWarning(message string) {
	l.log(WARN, "Msg: %s ", message)
}

func (l *Logger) LogError(message string) {
	l.log(ERROR, "Msg: %s ", message)
}

func (l *Logger) LogInfof(formatter string, data ...interface{}) {
	l.log(INFO, formatter, data...)
}

func (l *Logger) LogDebugf(formatter string, data ...interface{}) {
	l.log(DEBUG, formatter, data)
}

func (l *Logger) LogWarningf(formatter string, data ...interface{}) {
	l.log(WARN, formatter, data)
}

func (l *Logger) LogErrorf(formatter string, data ...interface{}) {
	l.log(ERROR, formatter, data)
}

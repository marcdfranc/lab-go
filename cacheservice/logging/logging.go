/*
	move formater to log
ok	remove fatal
	parameters should be interfaces
ok	move log implementation to private
	use FMT, errorFormat and FPrintf
ok	Remove the weberver
	ordenate the imports (default, external packages, my local packages)
ok	change the name of the Logger to Logging and the loggerIstance to logging
*/

package logging

import (
	"fmt"
	"io"
	"time"
)

type Logger interface {
	LogInfo(message string)
	LogDebug(message string)
	LogWarning(message string)
	LogError(message string)
	LogInfof(formatter string, data ...interface{})
	LogDebugf(formatter string, data ...interface{})
	LogWarningf(formatter string, data ...interface{})
	LogErrorf(formatter string, data ...interface{})
}

type logger struct {
	output io.Writer
	source string
}

func NewLogger(output io.Writer, source string) Logger {
	return &logger{
		output: output,
		source: source,
	}
}

func (l *logger) log(level LogLevel, template string, data ...any) {
	template = "[%s] Source: %s, Time: %s, " + template + "\n"
	args := []any{level.String(), l.source, time.Now().UTC().Format(time.RFC3339)}
	args = append(args, data...)
	logMessage := fmt.Sprintf(template, args...)
	l.output.Write([]byte(logMessage))
}

func (l *logger) LogInfo(message string) {
	l.log(INFO, "Msg: %s ", message)
}

func (l *logger) LogDebug(message string) {
	l.log(DEBUG, "Msg: %s ", message)
}

func (l *logger) LogWarning(message string) {
	l.log(WARN, "Msg: %s ", message)
}

func (l *logger) LogError(message string) {
	l.log(ERROR, "Msg: %s ", message)
}

func (l *logger) LogInfof(formatter string, data ...interface{}) {
	l.log(INFO, formatter, data...)
}

func (l *logger) LogDebugf(formatter string, data ...interface{}) {
	l.log(DEBUG, formatter, data)
}

func (l *logger) LogWarningf(formatter string, data ...interface{}) {
	l.log(WARN, formatter, data)
}

func (l *logger) LogErrorf(formatter string, data ...interface{}) {
	l.log(ERROR, formatter, data)
}

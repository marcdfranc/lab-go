// logger.go
package logger

import (
	"io"
	"os"
	"time"
)

type Logger struct {
	output    io.Writer
	formatter *LogFormatter
	Name      string
}

func NewLogger(output io.Writer, name string, formatter *LogFormatter) *Logger {
	return &Logger{
		output:    output,
		formatter: formatter,
		Name:      name,
	}
}

func (l *Logger) Log(level LogLevel, data map[string]interface{}) {
	// Adiciona campos padrão ao mapa de dados
	data["Name"] = l.Name
	data["Level"] = level.String()
	data["Time"] = time.Now().Format(time.RFC3339)

	logMessage := l.formatter.Format(data)
	l.output.Write([]byte(logMessage))

	if level == FATAL {
		os.Exit(1)
	}
}

func (l *Logger) LogInfo(data map[string]interface{}) {
	l.Log(INFO, data)
}

func (l *Logger) LogDebug(data map[string]interface{}) {
	l.Log(DEBUG, data)
}

func (l *Logger) LogWarning(data map[string]interface{}) {
	l.Log(WARN, data)
}

func (l *Logger) LogError(data map[string]interface{}) {
	l.Log(ERROR, data)
}

func (l *Logger) LogFatal(data map[string]interface{}) {
	l.Log(FATAL, data)
}

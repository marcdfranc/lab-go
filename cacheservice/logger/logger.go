package logger

type ILogger interface {
	Log(level LogLevel, message string, errors ...error)
}

type Logger struct {
	output IOutput
	Name   string
}

func NewLogger(output IOutput, name string) *Logger {
	return &Logger{
		output: output,
		Name:   name,
	}
}

func (l *Logger) Log(level LogLevel, message string, errors ...error) {
	l.output.Log(level, l.Name, message, errors...)
}

func (l *Logger) LogInfo(message string) {
	l.Log(INFO, message)
}

func (l *Logger) LogDebug(message string) {
	l.Log(DEBUG, message)
}

func (l *Logger) LogWarning(message string) {
	l.Log(WARN, message)
}

func (l *Logger) LogError(message string, errors ...error) {
	l.Log(ERROR, message, errors...)
}

func (l *Logger) LogFatal(message string, errors ...error) {
	l.Log(FATAL, message, errors...)
}

package logger

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNDEFINED"
	}
}

func (l *Logger) Log(message string) {
	l.log(UNDEFINED, message)
}

func (l *Logger) LogInfo(message string) {
	l.log(INFO, message)
}

func (l *Logger) LogDebug(message string) {
	l.log(DEBUG, message)
}

func (l *Logger) LogWarning(message string) {
	l.log(WARN, message)
}

func (l *Logger) LogError(message string, errors ...error) {
	l.log(ERROR, message, errors...)
}

func (l *Logger) LogFatal(message string, errors ...error) {
	l.log(FATAL, message, errors...)
}

package logger

// ILogger defines the logger interface.
type ILogger interface {
	Log(message string)
	LogInfo(message string)
	LogDebug(message string)
	LogWarning(message string)
	LogError(message string, errors ...error)
	LogFatal(message string, errors ...error)
}

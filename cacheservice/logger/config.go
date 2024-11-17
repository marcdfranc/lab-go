package logger

// LogLevel defines the level of logging.
type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	WARN
	ERROR
	FATAL
	UNDEFINED
)

// Logger represents the logger with a name.
type Logger struct {
	Name string
}

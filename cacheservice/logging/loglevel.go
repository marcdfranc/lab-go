// loglevel.go
package logging

type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	WARN
	ERROR
	FATAL
	UNDEFINED
)

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

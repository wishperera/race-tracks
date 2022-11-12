package log

const (
	TRACE = iota + 1
	DEBUG
	INFO
	WARN
	ERROR
	FATAL

	defaultLogLevel = INFO
)

type Logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	TraceF(formatString string, params ...interface{})
	DebugF(formatString string, params ...interface{})
	InfoF(formatString string, params ...interface{})
	WarnF(formatString string, params ...interface{})
	ErrorF(formatString string, params ...interface{})
	FatalF(formatString string, params ...interface{})
}

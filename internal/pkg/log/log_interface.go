package log

import (
	"fmt"
)

const (
	TRACE LogLevel = iota + 1
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var supportedLevelsLookupTable = map[string]LogLevel{
	"trace": TRACE,
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
}

// LogLevel : defines the log level for the application
type LogLevel int

// ParseLevelFromString : returns a LogLevel type by parsing the provided string
func ParseLevelFromString(lvl string) (l LogLevel, err error) {
	v, ok := supportedLevelsLookupTable[lvl]
	if !ok {
		return l, fmt.Errorf("unsupported log level: %s", lvl)
	}

	return v, nil
}

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

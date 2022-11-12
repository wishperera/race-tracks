package log

import (
	"os"

	"github.com/rs/zerolog"
)

type Log struct {
	logger zerolog.Logger
}

func NewLog(level LogLevel) *Log {
	l := &Log{}
	l.logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(l.mapLevelToZeroLogLevel(level))

	return l
}

func (s *Log) Trace(message string) {
	s.logger.Trace().Msg(message)
}

func (s *Log) Debug(message string) {
	s.logger.Debug().Msg(message)
}

func (s *Log) Info(message string) {
	s.logger.Info().Msg(message)
}

func (s *Log) Warn(message string) {
	s.logger.Warn().Msg(message)
}

func (s *Log) Error(message string) {
	s.logger.Error().Msg(message)
}

func (s *Log) Fatal(message string) {
	s.logger.Fatal().Msg(message)
}

func (s *Log) TraceF(formatString string, params ...interface{}) {
	s.logger.Trace().Msgf(formatString, params...)
}

func (s *Log) DebugF(formatString string, params ...interface{}) {
	s.logger.Debug().Msgf(formatString, params...)
}

func (s *Log) InfoF(formatString string, params ...interface{}) {
	s.logger.Info().Msgf(formatString, params...)
}

func (s *Log) WarnF(formatString string, params ...interface{}) {
	s.logger.Warn().Msgf(formatString, params...)
}

func (s *Log) ErrorF(formatString string, params ...interface{}) {
	s.logger.Error().Msgf(formatString, params...)
}

func (s *Log) FatalF(formatString string, params ...interface{}) {
	s.logger.Fatal().Msgf(formatString, params...)
}

// mapLevelToZeroLogLevel : maps the log level to zero log compatible level value
func (s *Log) mapLevelToZeroLogLevel(lvl LogLevel) zerolog.Level {
	switch lvl {
	case TRACE:
		return zerolog.TraceLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case WARN:
		return zerolog.WarnLevel
	case ERROR:
		return zerolog.ErrorLevel
	case FATAL:
		return zerolog.FatalLevel
	default:
		return zerolog.NoLevel
	}
}

package log

import (
	"github.com/rs/zerolog"
	"os"
)

type StandardLog struct {
	logger zerolog.Logger
}

func NewLog() *StandardLog {
	return &StandardLog{
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

func (s *StandardLog) Trace(message string) {
	s.logger.Trace().Msg(message)
}

func (s *StandardLog) Debug(message string) {
	s.logger.Debug().Msg(message)
}

func (s *StandardLog) Info(message string) {
	s.logger.Info().Msg(message)
}

func (s *StandardLog) Warn(message string) {
	s.logger.Warn().Msg(message)
}

func (s *StandardLog) Error(message string) {
	s.logger.Error().Msg(message)
}

func (s *StandardLog) Fatal(message string) {
	s.logger.Fatal().Msg(message)
}

func (s *StandardLog) TraceF(formatString string, params ...interface{}) {
	s.logger.Trace().Msgf(formatString, params...)
}

func (s *StandardLog) DebugF(formatString string, params ...interface{}) {
	s.logger.Debug().Msgf(formatString, params...)
}

func (s *StandardLog) InfoF(formatString string, params ...interface{}) {
	s.logger.Info().Msgf(formatString, params...)
}

func (s *StandardLog) WarnF(formatString string, params ...interface{}) {
	s.logger.Warn().Msgf(formatString, params...)
}

func (s *StandardLog) ErrorF(formatString string, params ...interface{}) {
	s.logger.Error().Msgf(formatString, params...)
}

func (s *StandardLog) FatalF(formatString string, params ...interface{}) {
	//TODO implement me
	panic("implement me")
}

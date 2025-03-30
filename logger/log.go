package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strings"
)

func Log() {

	globalLog()

}
func globalLog() {
	globalLevel := loglevel("debug")
	zerolog.SetGlobalLevel(globalLevel)
	log.Info().Msgf("Logging level set to %s", "debug")
}

func loglevel(level string) zerolog.Level {
	level = strings.ToUpper(level)
	switch level {
	case "FATAL":
		return zerolog.FatalLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "INFO":
		return zerolog.InfoLevel
	case "DEBUG":
		return zerolog.DebugLevel
	default:
		log.Warn().Msgf("defult level set  %s", "INFO")
		return zerolog.InfoLevel
	}
}

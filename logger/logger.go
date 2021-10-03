package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

type StandardLogger struct {
	*log.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = log.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &log.JSONFormatter{}
	standardLogger.Logger.SetOutput(os.Stdout)
	standardLogger.Logger.SetFormatter(&log.JSONFormatter{})

	// TODO: - Env setup
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.DebugLevel
	}

	standardLogger.Logger.SetLevel(logLevel)
	return standardLogger
}
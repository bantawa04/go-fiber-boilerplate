package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

// SetUpLogger settings
func SetUpLogger(env *Env) *Logger {
	logger := &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if env.APP_DEBUG {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.WithField("app_name", env.APP_NAME).Info("Logger initialized")
	return logger
}

func GetLogger() *Logger {
	return logger
}

package logger

import (
	logger "github.com/sirupsen/logrus"
)

func NewLogger() *logger.Logger {
	return logger.New()
}

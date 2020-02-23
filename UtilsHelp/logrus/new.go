package logrus

import "github.com/sirupsen/logrus"

func NewLog(level uint32) *logrus.Logger {
	logger := logrus.New()
	return logger
}

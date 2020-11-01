package logrus

import (
	"github.com/sirupsen/logrus"
)

var logging *logrus.Logger

func InitLog(logPath string, maxAge int, compress bool) *logrus.Logger {
	logging = NewLoggers(logPath, maxAge, compress)
	return logging
}

func GetLog() *logrus.Logger {
	return logging
}

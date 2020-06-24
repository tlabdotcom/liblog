package liblog

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func getLogLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}

func getFormatter() logrus.Formatter {
	formatter := &logrus.TextFormatter{}
	formatter.FullTimestamp = true
	return formatter
}

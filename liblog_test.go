package liblog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	var print logrus.FieldLogger
	c := ConfigLogger{
		Stdout:     true,
		Level:      "DEBUG",
		OutputFile: "log.text",
	}
	print = c.NewLogger()
	print.Infof("Hi")
}

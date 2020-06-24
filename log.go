package liblog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type (
	// ConfigLogger config logger.
	ConfigLogger struct {
		// Stdout is true if the output needs to goto standard out
		Stdout bool `yaml:"stdout"`
		// Level is the desired log level
		Level string `yaml:"level"`
		// OutputFile is the path to the log output file
		OutputFile string `yaml:"outputFile"`
	}
)

const fileMode = os.FileMode(0644)

// NewLogger is used to initial Logger
func (l *ConfigLogger) NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	logger.Level = getLogLevel(l.Level)
	logger.Formatter = getFormatter()

	if l.Stdout {
		logger.Out = os.Stdout
	}

	if len(l.OutputFile) > 0 {
		outFile := createLogFile(l.OutputFile)
		logger.Out = outFile
		if l.Stdout {
			logger.Out = io.MultiWriter(os.Stdout, outFile)
		}
	}
	return logger
}

func createLogFile(path string) *os.File {
	dir := filepath.Dir(path)
	if len(dir) > 0 && dir != "." {
		if err := os.MkdirAll(dir, fileMode); err != nil {
			log.Fatalf("error creating log directory %v, err=%v", dir, err)
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, fileMode)
	if err != nil {
		log.Fatalf("error creating log file %v, err=%v", path, err)
	}
	return file
}

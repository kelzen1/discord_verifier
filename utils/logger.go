package utils

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

var (
	once    sync.Once
	logger  *logrus.Logger
	logFile *os.File
)

func Logger() *logrus.Logger {
	once.Do(func() {

		logger = logrus.New()

		f, err := os.OpenFile("verifier.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			logger.Panicln("error opening file:", err)
		}

		logFile = f

		writer := io.MultiWriter(f, os.Stdout)
		logger.SetOutput(writer)

	})

	return logger
}

func ShutdownLogger() {
	_ = logFile.Close()
}

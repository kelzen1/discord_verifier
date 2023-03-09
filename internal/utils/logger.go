package utils

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"sync"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func Logger() *logrus.Logger {
	once.Do(func() {

		logger = logrus.New()
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})

		f, err := os.OpenFile(path.Join("data", "verifier.log"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			logger.Panicln("error opening file:", err)
		}

		writer := io.MultiWriter(f, os.Stdout)
		logger.SetOutput(writer)

	})

	return logger
}

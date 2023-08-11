package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func NewDefaultLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "ts",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
		},
		TimestampFormat: time.RFC3339,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	return log
}

func SetLogName(name string, logger *logrus.Logger) *logrus.Entry {
	return logger.WithField("logger", name)
}

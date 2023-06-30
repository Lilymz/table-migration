package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

var DaoLog = logrus.New()

func init() {
	DaoLog.Level = logrus.InfoLevel
	DaoLog.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	DaoLog.SetOutput(os.Stdout)

}

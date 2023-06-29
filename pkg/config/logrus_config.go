package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var DaoLog = logrus.New()

func init() {
	DaoLog.Level = logrus.InfoLevel
	DaoLog.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	DaoLog.SetOutput(os.Stdout)
	handler := func() {
		fmt.Println("出现了致命的FATAL错误！执行关闭前逻辑")
		// A common pattern is to re-use fields between logging statements by re-using
		// the logrus.Entry returned from WithFields()
		contextLogger := DaoLog.WithFields(logrus.Fields{
			"common": "this is a common field",
			"other":  "I also should be logged always",
		})

		contextLogger.Logger.Warn("I'll be logged with common and other field")
		contextLogger.Info("Me too")
	}
	logrus.RegisterExitHandler(handler)
}

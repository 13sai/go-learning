package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	l, _ := logrus.ParseLevel("info")
	logrus.SetLevel(l)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)

	logrus.Debug("debug", "what")
	logrus.Warn("warn", "what")
	logrus.Info("warn", "what")
}
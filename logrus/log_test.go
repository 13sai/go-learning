package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
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

package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	l, _ := logrus.ParseLevel("info")
	logrus.SetLevel(l)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(true)

	logrus.Debug("debug", "what")
	logrus.Warn("warn", "what")
	logrus.Info("warn", "what")

	logrus.WithFields(logrus.Fields{
		"err": "err",
	}).Info("jshfjsh")
}

func TestLogWithFields(t *testing.T) {
	log := logrus.WithFields(logrus.Fields{
		"app_id": 1,
	})

	log.Info("hello")
	log.Errorf("wrong %s", "what")
}

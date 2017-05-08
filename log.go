package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLogger() {
	Logger = logrus.New()
	Logger.Level = logrus.DebugLevel
	Logger.Out = os.Stdout
}

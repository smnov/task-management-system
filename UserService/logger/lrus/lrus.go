package lrus

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	return logrus.New()
}
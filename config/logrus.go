package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *Config) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(6))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

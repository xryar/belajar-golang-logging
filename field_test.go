package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "andaru").Info("Hello World")

	logger.WithField("username", "arya").
		WithField("name", "Arya Rizki").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "arya",
		"name":     "Arya Rizki",
	}).Info("Hello World")
}

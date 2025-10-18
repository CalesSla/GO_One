package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in")

}

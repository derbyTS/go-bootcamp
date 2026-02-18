package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("This is infor")
	log.Warn("This is warn")
	log.Error("This is error")

	log.WithFields(logrus.Fields{
		"userName": "John Doe",
		"method":   "GET",
	}).Info("User logged in")
}

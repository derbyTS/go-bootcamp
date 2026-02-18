package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println(err)
	}

	defer logger.Sync()

	logger.Info("This is a message")
	logger.Info("User logged in", zap.String("userName", "John Doe"), zap.String("method", "GET"))
}

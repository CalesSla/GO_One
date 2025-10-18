package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error initializing Zap logger:", err)
		return
	}

	defer logger.Sync()

	logger.Info("This is an info message")
	logger.Info("User logged in", zap.String("username", "John Doe"), zap.String("method", "GET"))

}

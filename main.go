package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger()
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize config: %v", err)
		panic(err)
	}
	router.Initialize("8080")
}

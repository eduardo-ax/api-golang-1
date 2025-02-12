package main

import (
	config "github.com/eduardo-ax/api-golang-1/config/database"
	"github.com/eduardo-ax/api-golang-1/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}

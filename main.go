package main

import (
	"github.com/alexduzi/job-openings/config"
	"github.com/alexduzi/job-openings/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}

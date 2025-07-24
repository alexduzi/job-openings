package main

import (
	"github.com/alexduzi/job-openings/config"
	"github.com/alexduzi/job-openings/router"
)

var (
	logger *config.Logger
)

// @title           Jog openings
// @version         1.0
// @description     This is a sample server for manage the open positions with a full CRUD operations
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}

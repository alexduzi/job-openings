package router

import (
	docs "github.com/alexduzi/job-openings/docs"
	"github.com/alexduzi/job-openings/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)

		v1.GET("/opening/:id", handler.GetOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}

}

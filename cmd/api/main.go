package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/cmd/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	CategoryRoutes(router)

	router.Run(":8080")
}

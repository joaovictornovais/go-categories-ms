package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/cmd/api/controllers"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})
	categoryRoutes.GET("", func(ctx *gin.Context) {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})
	categoryRoutes.DELETE("", func(ctx *gin.Context) {
		controllers.DeleteCategory(ctx, inMemoryCategoryRepository)
	})
	categoryRoutes.PUT("", func(ctx *gin.Context) {
		controllers.UpdateCategory(ctx, inMemoryCategoryRepository)
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
	use_cases "github.com/joaovictornovais/go-categories-ms/internal/use-cases"
)

type listCategoriesInput struct {
	Name string `json:"name"`
}

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body listCategoriesInput

	_ = ctx.ShouldBindJSON(&body)

	if body.Name != "" {
		useCase := use_cases.NewGetCategoryUseCase(repository)
		category, err := useCase.Execute(body.Name)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success":  true,
			"category": category,
		})
		return
	}

	useCase := use_cases.NewListCategoriesUseCase(repository)
	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":    true,
		"categories": categories,
	})
}

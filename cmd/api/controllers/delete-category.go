package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
	use_cases "github.com/joaovictornovais/go-categories-ms/internal/use-cases"
)

type deleteCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func DeleteCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body deleteCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	use_case := use_cases.NewDeleteCategoryUseCase(repository)

	err := use_case.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(204, gin.H{})

}

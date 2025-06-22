package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
	use_cases "github.com/joaovictornovais/go-categories-ms/internal/use-cases"
)

type updateCategoryInput struct {
	CurrentName string `json:"currentName" binding:"required"`
	NewName     string `json:"newName" binding:"required"`
}

// @Summary Update category
// @Schemes http
// @Description Update category from database. If category 'currentName' provided in body exists and 'newName' don't, update category name with Success. If not, throws error.
// @Tags Category
// @Accept json
// @Produce json
// @Param request body updateCategoryInput true "Category new data"
// @Success 200 {array} entities.Category
// @Failure 404 {object} string "Category not found"
// @Router /categories [put]
func UpdateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body updateCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := use_cases.NewUpdateCategoryUseCase(repository)

	category, err := useCase.Execute(body.CurrentName, body.NewName)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"category": category,
	})
}

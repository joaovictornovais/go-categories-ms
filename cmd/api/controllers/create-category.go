package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
	use_cases "github.com/joaovictornovais/go-categories-ms/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

// @BasePath categories

// @Summary Create category
// @Schemes
// @Description Persist category on database
// @Tags Category
// @Accept json
// @Produce json
// @Param request body createCategoryInput true "Category data"
// @Success 201 {object} entities.Category
// @Router /categories [post]
func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	useCase := use_cases.NewCreateCategoryUseCase(repository)

	category, err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"category": category,
	})
}

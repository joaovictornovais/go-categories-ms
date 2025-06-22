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

// @Summary Get categories
// @Schemes http
// @Description Get categories from database. If 'name' is provided in body, returns a single category. If not, returns all.
// @Tags Category
// @Accept json
// @Produce json
// @Param request body listCategoriesInput false "Category filter"
// @Success 200 {array} entities.Category
// @Failure 404 {object} string "Category not found"
// @Router /categories [get]
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

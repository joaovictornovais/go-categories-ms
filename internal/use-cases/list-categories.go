package use_cases

import (
	"github.com/joaovictornovais/go-categories-ms/internal/entities"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
)

type listCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUseCase(repository repositories.ICategoryRepository) *listCategoriesUseCase {
	return &listCategoriesUseCase{
		repository: repository,
	}
}

func (u *listCategoriesUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

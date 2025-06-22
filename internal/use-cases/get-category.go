package use_cases

import (
	"github.com/joaovictornovais/go-categories-ms/internal/entities"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
)

type getCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewGetCategoryUseCase(repository repositories.ICategoryRepository) *getCategoryUseCase {
	return &getCategoryUseCase{
		repository: repository,
	}
}

func (u *getCategoryUseCase) Execute(name string) (*entities.Category, error) {
	category, err := u.repository.Get(name)

	if err != nil {
		return nil, err
	}

	return category, nil
}

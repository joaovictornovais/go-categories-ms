package use_cases

import (
	"github.com/joaovictornovais/go-categories-ms/internal/entities"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
)

type updateCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewUpdateCategoryUseCase(repository repositories.ICategoryRepository) *updateCategoryUseCase {
	return &updateCategoryUseCase{
		repository: repository,
	}
}

func (u *updateCategoryUseCase) Execute(currentName string, newName string) (*entities.Category, error) {
	category, err := u.repository.Update(currentName, newName)

	if err != nil {
		return nil, err
	}

	return category, nil
}

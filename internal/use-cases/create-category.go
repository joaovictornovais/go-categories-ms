package use_cases

import (
	"log"

	"github.com/joaovictornovais/go-categories-ms/internal/entities"
	"github.com/joaovictornovais/go-categories-ms/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil

}

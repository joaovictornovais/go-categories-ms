package use_cases

import "github.com/joaovictornovais/go-categories-ms/internal/repositories"

type deleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{
		repository: repository,
	}
}

func (u *deleteCategoryUseCase) Execute(name string) error {
	err := u.repository.Delete(name)

	if err != nil {
		return err
	}

	return nil
}

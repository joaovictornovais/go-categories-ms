package repositories

import "github.com/joaovictornovais/go-categories-ms/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
	Get(name string) (*entities.Category, error)
	Delete(name string) error
	Update(currentName string, newName string) (*entities.Category, error)
}

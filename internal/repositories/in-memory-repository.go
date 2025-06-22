package repositories

import (
	"fmt"
	"time"

	"github.com/joaovictornovais/go-categories-ms/internal/entities"
)

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	_, err := r.Get(category.Name)

	if err != nil {
		r.db = append(r.db, category)
		return nil
	}

	return fmt.Errorf("already exists a category with name '%s'", category.Name)

}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}

func (r *inMemoryCategoryRepository) Get(name string) (*entities.Category, error) {
	var category *entities.Category

	for i := range r.db {
		if r.db[i].Name == name {
			category = r.db[i]
			break
		}
	}

	if category != nil {
		return category, nil
	}

	return nil, fmt.Errorf("category with name %s not found", name)
}

func (r *inMemoryCategoryRepository) Delete(name string) error {
	var newDb []*entities.Category

	for i := range r.db {
		if r.db[i].Name != name {
			newDb = append(newDb, r.db[i])
		}
	}

	if newDb == nil {
		r.db = make([]*entities.Category, 0)
	} else {
		r.db = newDb
	}

	return nil
}

func (r *inMemoryCategoryRepository) Update(currentName string, newName string) (*entities.Category, error) {

	_, err := r.Get(newName)

	if err != nil {
		for i := range r.db {
			if r.db[i].Name == currentName {
				r.db[i].Name = newName
				r.db[i].UpdatedAt = time.Now()
				return r.db[i], nil
			}
		}
		return nil, fmt.Errorf("there is no category with name '%s' on database", currentName)
	}

	return nil, fmt.Errorf("already exists a category with name '%s'", newName)
}

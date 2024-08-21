package repositories

import "github.com/gabriel-vidile/go-microservice/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}

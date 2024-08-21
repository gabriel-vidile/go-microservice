package services

import (
	"github.com/gabriel-vidile/go-microservice/internal/entities"
	"github.com/gabriel-vidile/go-microservice/internal/repositories"
)

type getCategoryService struct {
	repository repositories.ICategoryRepository
}

func NewGetCategoryService(repository repositories.ICategoryRepository) *getCategoryService {
	return &getCategoryService{repository}
}

func (u *getCategoryService) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

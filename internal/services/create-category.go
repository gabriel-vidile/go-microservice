package services

import (
	"github.com/gabriel-vidile/go-microservice/internal/entities"
	"github.com/gabriel-vidile/go-microservice/internal/repositories"
)

type createCategoryService struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryService(repository repositories.ICategoryRepository) *createCategoryService {
	return &createCategoryService{repository}
}

func (u *createCategoryService) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}
	u.repository.Save(category)
	return nil
}

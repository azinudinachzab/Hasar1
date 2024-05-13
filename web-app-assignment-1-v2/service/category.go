package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category *model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]*model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	err := c.categoryRepository.Store(category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryService) Update(id int, category *model.Category) error {
	// Implement the Update function here
	return nil
}

func (c *categoryService) Delete(id int) error {
	// Implement the Delete function here
	return nil
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	// Implement the GetByID function here
	return nil, nil
}

func (c *categoryService) GetList() ([]*model.Category, error) {
	// Implement the GetList function here
	return nil, nil
}

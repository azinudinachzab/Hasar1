package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type CategoryRepository interface {
	Store(category *model.Category) error
	Update(id int, category *model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]*model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(category *model.Category) error {
	c.filebasedDb.StoreCategory(*category)

	return nil
}

func (c *categoryRepository) Update(id int, category *model.Category) error {
	//beginanswer
	c.filebasedDb.UpdateCategory(id, *category)

	return nil
}

func (c *categoryRepository) Delete(id int) error {
	// Implement the Delete function here
	c.filebasedDb.DeleteCategory(id)

	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	// Implement the GetByID function here

	return nil, nil
}

func (c *categoryRepository) GetList() ([]*model.Category, error) {
	// Implement the GetList function here
	return nil, nil
}

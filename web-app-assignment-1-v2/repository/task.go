package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]*model.Task, error)
	GetTaskCategory(id int) ([]*model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)
	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	// Implement the Update function here
	return nil
}

func (t *taskRepository) Delete(id int) error {
	// Implement the Delete function here
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	// Implement the GetByID function here
	return nil, nil
}

func (t *taskRepository) GetList() ([]*model.Task, error) {
	// Implement the GetList function here
	return nil, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]*model.TaskCategory, error) {
	// Implement the GetTaskCategory function here
	return nil, nil
}

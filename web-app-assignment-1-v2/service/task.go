package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]*model.Task, error)
	GetTaskCategory(id int) ([]*model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (s *taskService) Store(task *model.Task) error {
	err := s.taskRepository.Store(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Update(task *model.Task) error {
	// Implement the Update function here
	return nil
}

func (s *taskService) Delete(id int) error {
	// Implement the Delete function here
	return nil
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	// Implement the GetByID function here
	return nil, nil
}

func (s *taskService) GetList() ([]*model.Task, error) {
	// Implement the GetList function here
	return nil, nil
}

func (s *taskService) GetTaskCategory(id int) ([]*model.TaskCategory, error) {
	// Implement the GetTaskCategory function here
	return nil, nil
}

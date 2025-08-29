package services

import (
	"my-go-api/internal/models"
	"my-go-api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() []models.Task {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return []models.Task{}
	}
	return tasks
}

func (s *TaskService) AddTask(newTask models.Task) models.Task {
	task, err := s.repo.AddTask(newTask)
	if err != nil {
		return models.Task{}
	}
	return task
}

func (s *TaskService) UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	return s.repo.UpdateTask(id, updatedTask)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}

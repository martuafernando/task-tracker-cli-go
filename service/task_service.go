package service

import (
	"task-tracker-cli/models"
	"task-tracker-cli/repository"
)

type TaskService struct {
	Repository *repository.TaskRepository
}

func (r TaskService) AddTask(taskname string) error {
	task := models.Task{
		Name:   taskname,
		Status: models.Todo,
	}
	return r.Repository.Create(task)
}

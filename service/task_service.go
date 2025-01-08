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

func (r TaskService) UpdateTask(id int, taskname string) error {
	udpateTask := models.Task{
		Name: taskname,
	}

	return r.Repository.Update(id, udpateTask)
}

package service

import (
	"fmt"
	"task-tracker-cli/models"
	"task-tracker-cli/repository"
)

type TaskService struct {
	Repository *repository.TaskRepository
}

func (r *TaskService) AddTask(taskname string) error {
	task := models.Task{
		Name:   taskname,
		Status: models.Todo,
	}
	return r.Repository.Create(task)
}

func (r *TaskService) UpdateTask(id int, taskname string) error {
	udpateTask := models.Task{
		Name: taskname,
	}

	return r.Repository.Update(id, udpateTask)
}

func (r *TaskService) DeleteTask(id int) error {
	_, err := r.Repository.Get(id)

	if err != nil {
		return fmt.Errorf("task not found")
	}

	return r.Repository.Delete(id)
}

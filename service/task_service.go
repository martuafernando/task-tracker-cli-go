package service

import (
	"fmt"
	"task-tracker-cli/models"
	"task-tracker-cli/repository"
)

type TaskService interface {
	AddTask(taskname string) error
	UpdateTask(id int, task models.Task) error
	DeleteTask(id int) error
	GetAllTask(status *models.Status) []models.Task
}

type TaskServiceImpl struct {
	Repository repository.TaskRepository
}

func (r *TaskServiceImpl) AddTask(taskname string) error {
	task := models.Task{
		Name:   taskname,
		Status: models.Todo,
	}
	return r.Repository.Create(task)
}

func (r *TaskServiceImpl) UpdateTask(id int, task models.Task) error {
	return r.Repository.Update(id, task)
}

func (r *TaskServiceImpl) DeleteTask(id int) error {
	_, err := r.Repository.Get(id)

	if err != nil {
		return fmt.Errorf("task not found")
	}

	return r.Repository.Delete(id)
}

func (r *TaskServiceImpl) GetAllTask(status *models.Status) []models.Task {
	return r.Repository.GetAll(status)
}

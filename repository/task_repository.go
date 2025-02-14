package repository

import (
	"fmt"
	"task-tracker-cli/models"
	"task-tracker-cli/storage"
)

type TaskRepository interface {
	Create(task models.Task) error
	Update(id int, task models.Task) error
	Delete(id int) error
	Get(id int) (models.Task, error)
	GetAll(status *models.Status) []models.Task
}

type TaskRepositoryImpl struct {
	Filestorage storage.FileStorage
}

func (r *TaskRepositoryImpl) Create(task models.Task) error {
	var tasks []models.Task

	err := r.Filestorage.Load(&tasks)

	if err != nil {
		return fmt.Errorf("failed to load from file: %w", err)
	}

	task.Id = getNewId(tasks)
	tasks = append(tasks, task)
	return r.Filestorage.Save(tasks)
}

func (r *TaskRepositoryImpl) Update(id int, task models.Task) error {
	var tasks []models.Task

	err := r.Filestorage.Load(&tasks)

	if err != nil {
		return fmt.Errorf("failed to load from file: %w", err)
	}

	currentTask := getById(tasks, id)

	if currentTask == nil {
		return fmt.Errorf("task with id %d is not found", id)
	}

	if task.Name != "" {
		currentTask.Name = task.Name
	}

	if task.Status != 0 {
		currentTask.Status = task.Status
	}

	return r.Filestorage.Save(tasks)
}

func (r *TaskRepositoryImpl) Delete(id int) error {
	var tasks []models.Task

	err := r.Filestorage.Load(&tasks)

	if err != nil {
		return fmt.Errorf("failed to load from file: %w", err)
	}

	var index = 0
	for i := range tasks {
		if tasks[i].Id == id {
			index = i
		}
	}

	return r.Filestorage.Save(append(tasks[:index], tasks[index+1:]...))
}

func (r *TaskRepositoryImpl) Get(id int) (models.Task, error) {
	var tasks []models.Task

	err := r.Filestorage.Load(&tasks)

	if err != nil {
		return models.Task{}, fmt.Errorf("failed to load from file: %w", err)
	}

	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return models.Task{}, fmt.Errorf("task with id %d not found", id)
}

func (r *TaskRepositoryImpl) GetAll(status *models.Status) []models.Task {
	var tasks []models.Task

	err := r.Filestorage.Load(&tasks)

	if err != nil {
		return []models.Task{}
	}

	if status == nil {
		return tasks
	}

	filteredTask := []models.Task{}
	for i := range tasks {
		if tasks[i].Status == *status {
			filteredTask = append(filteredTask, tasks[i])
		}
	}

	return filteredTask
}

func getNewId(tasks []models.Task) int {
	maxId := 0
	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	return maxId + 1
}

func getById(tasks []models.Task, id int) *models.Task {
	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i]
		}
	}

	return nil
}

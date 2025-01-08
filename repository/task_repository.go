package repository

import (
	"fmt"
	"task-tracker-cli/models"
	"task-tracker-cli/storage"
)

type TaskRepository struct {
	Filename string
}

func (r *TaskRepository) Create(task models.Task) error {
	var tasks []models.Task

	err := storage.LoadFromFile(r.Filename, &tasks)

	if err != nil {
		return fmt.Errorf("failed to load from file: %w", err)
	}

	task.Id = getNewId(tasks)
	tasks = append(tasks, task)
	return storage.SaveToFile(r.Filename, tasks)
}

func (r *TaskRepository) Update(id int, task models.Task) error {
	var tasks []models.Task

	err := storage.LoadFromFile(r.Filename, &tasks)

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

	return storage.SaveToFile(r.Filename, tasks)
}

func (r *TaskRepository) Delete(id int) error {
	var tasks []models.Task

	err := storage.LoadFromFile(r.Filename, &tasks)

	if err != nil {
		return fmt.Errorf("failed to load from file: %w", err)
	}

	var index = 0
	for i := range tasks {
		if tasks[i].Id == id {
			index = i
		}
	}

	return storage.SaveToFile(r.Filename, append(tasks[:index], tasks[index+1:]...))
}

func (r *TaskRepository) Get(id int) (models.Task, error) {
	var tasks []models.Task

	err := storage.LoadFromFile(r.Filename, &tasks)

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

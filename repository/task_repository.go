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

func Edit(task models.Task) error {
	return nil
}

func Delete(id int) error {
	return nil
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

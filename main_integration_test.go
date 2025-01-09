package main

import (
	"os"
	"os/exec"
	"task-tracker-cli/models"
	"task-tracker-cli/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

const FILENAME = "./data/data.json"

func TestMain(m *testing.M) {
	os.Remove(FILENAME)

	m.Run()
}

func TestMainIntegration(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".")

	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	assert.Error(t, err)
}

func TestMainIntegrationAddTask(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "add", "testing task")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	tasks := []models.Task{}
	fileStorage.Load(&tasks)

	assert.NoError(t, err)
	assert.Equal(t, []models.Task{{Id: 1, Name: "testing task", Status: models.Todo}}, tasks)
}

func TestMainIntegrationUpdateTask(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "update", "1", "testing task")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{{
		Id:     1,
		Name:   "task initial",
		Status: models.Todo,
	}})

	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	tasks := []models.Task{}
	fileStorage.Load(&tasks)

	assert.NoError(t, err)
	assert.Equal(t, []models.Task{{Id: 1, Name: "testing task", Status: models.Todo}}, tasks)
}

func TestMainIntegrationDeleteTask(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "delete", "1")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{{
		Id:     1,
		Name:   "task initial",
		Status: models.Todo,
	}})

	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	tasks := []models.Task{}
	fileStorage.Load(&tasks)

	assert.NoError(t, err)
	assert.Equal(t, []models.Task{}, tasks)
}

func TestMainIntegrationMarkInProgress(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "mark-in-progress", "1")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{{
		Id:     1,
		Name:   "task initial",
		Status: models.Todo,
	}})

	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	tasks := []models.Task{}
	fileStorage.Load(&tasks)

	assert.NoError(t, err)
	assert.Equal(t, []models.Task{{Id: 1, Name: "task initial", Status: models.Inprogress}}, tasks)
}

func TestMainIntegrationDone(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "mark-done", "1")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{{
		Id:     1,
		Name:   "task initial",
		Status: models.Todo,
	}})

	// Act
	_, err := cmd.CombinedOutput()

	// Assert
	tasks := []models.Task{}
	fileStorage.Load(&tasks)

	assert.NoError(t, err)
	assert.Equal(t, []models.Task{{Id: 1, Name: "task initial", Status: models.Done}}, tasks)
}

func TestMainIntegrationList(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "list")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{
		{Id: 1, Name: "task todo", Status: models.Todo},
		{Id: 2, Name: "task in-progress", Status: models.Inprogress},
		{Id: 3, Name: "task done", Status: models.Done},
	})

	// Act
	output, err := cmd.CombinedOutput()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "1 task todo (Todo)\n2 task in-progress (In Progress)\n3 task done (Done)\n", string(output))
}

func TestMainIntegrationListTodo(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "list", "todo")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{
		{Id: 1, Name: "task todo", Status: models.Todo},
		{Id: 2, Name: "task in-progress", Status: models.Inprogress},
		{Id: 3, Name: "task done", Status: models.Done},
	})

	// Act
	output, err := cmd.CombinedOutput()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "1 task todo (Todo)\n", string(output))
}

func TestMainIntegrationListInProgress(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "list", "in-progress")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{
		{Id: 1, Name: "task todo", Status: models.Todo},
		{Id: 2, Name: "task in-progress", Status: models.Inprogress},
		{Id: 3, Name: "task done", Status: models.Done},
	})

	// Act
	output, err := cmd.CombinedOutput()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "2 task in-progress (In Progress)\n", string(output))
}

func TestMainIntegrationListDone(t *testing.T) {
	// Arrange
	cmd := exec.Command("go", "run", ".", "list", "done")
	fileStorage := storage.FileStorageImpl{
		Filename: FILENAME,
	}
	fileStorage.Save([]models.Task{
		{Id: 1, Name: "task todo", Status: models.Todo},
		{Id: 2, Name: "task in-progress", Status: models.Inprogress},
		{Id: 3, Name: "task done", Status: models.Done},
	})

	// Act
	output, err := cmd.CombinedOutput()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "3 task done (Done)\n", string(output))
}

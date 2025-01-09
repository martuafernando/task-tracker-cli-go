package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"task-tracker-cli/models"
	"task-tracker-cli/repository"
	"testing"
)

func TestTaskService_AddTask(t *testing.T) {
	// Arrange
	taskRepositoryMock := new(repository.TaskRepositoryMock)
	taskRepositoryMock.On("Create", mock.AnythingOfType("Task")).Return(nil)
	taskService := TaskServiceImpl{taskRepositoryMock}

	// Act
	err := taskService.AddTask("testing")

	// Assert
	assert.NoError(t, err)
	taskRepositoryMock.AssertNumberOfCalls(t, "Create", 1)
	taskRepositoryMock.AssertCalled(t, "Create", models.Task{
		Name:   "testing",
		Status: models.Todo,
	})
}

func TestTaskService_UpdateTask(t *testing.T) {
	// Arrange
	taskRepositoryMock := new(repository.TaskRepositoryMock)
	taskRepositoryMock.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("Task")).Return(nil)
	taskService := TaskServiceImpl{taskRepositoryMock}
	task := models.Task{
		Id:     1,
		Name:   "testing",
		Status: models.Todo,
	}

	// Act
	err := taskService.UpdateTask(1, task)

	// Assert
	assert.NoError(t, err)
	taskRepositoryMock.AssertNumberOfCalls(t, "Update", 1)
	taskRepositoryMock.AssertCalled(t, "Update", 1, models.Task{
		Id:     1,
		Name:   "testing",
		Status: models.Todo,
	})
}

func TestTaskService_DeleteTask(t *testing.T) {
	// Arrange
	taskRepositoryMock := new(repository.TaskRepositoryMock)
	taskRepositoryMock.On("Get", mock.AnythingOfType("int")).Return(models.Task{
		Id:     0,
		Name:   "testing",
		Status: models.Todo,
	}, nil)
	taskRepositoryMock.On("Delete", mock.AnythingOfType("int")).Return(nil)
	taskService := TaskServiceImpl{taskRepositoryMock}

	// Act
	err := taskService.DeleteTask(1)

	// Assert
	assert.NoError(t, err)
	taskRepositoryMock.AssertNumberOfCalls(t, "Delete", 1)
	taskRepositoryMock.AssertCalled(t, "Delete", 1)
}

func TestTaskService_GetAllTask(t *testing.T) {
	// Arrange
	taskRepositoryMock := new(repository.TaskRepositoryMock)
	taskRepositoryMock.On("GetAll", mock.Anything).Return([]models.Task{
		{Id: 1, Name: "testing", Status: models.Todo},
	}, nil)
	taskService := TaskServiceImpl{taskRepositoryMock}

	// Act
	tasks := taskService.GetAllTask(nil)

	// Assert
	assert.Len(t, tasks, 1)
	assert.Equal(t, tasks, []models.Task{
		{Id: 1, Name: "testing", Status: models.Todo},
	})
	taskRepositoryMock.AssertNumberOfCalls(t, "GetAll", 1)
}

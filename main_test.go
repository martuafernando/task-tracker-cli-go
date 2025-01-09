package main

import (
	"github.com/stretchr/testify/mock"
	"task-tracker-cli/models"
	"task-tracker-cli/service"
	"testing"
)

func TestAddTask(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("AddTask", mock.AnythingOfType("string")).Return(nil)

	// Act
	addTask(taskServiceMock, []string{"add", "testing task"})

	// Assert
	taskServiceMock.AssertCalled(t, "AddTask", "testing task")
}

func TestUpdateTask(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("UpdateTask", mock.AnythingOfType("int"), mock.AnythingOfType("Task")).Return(nil)
	task := models.Task{
		Name: "testing task",
	}
	// Act
	updateTask(taskServiceMock, []string{"update", "1", "testing task"})

	// Assert
	taskServiceMock.AssertCalled(t, "UpdateTask", 1, task)
}

func TestDeleteTask(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("DeleteTask", mock.AnythingOfType("int")).Return(nil)

	// Act
	deleteTask(taskServiceMock, []string{"update", "1"})

	// Assert
	taskServiceMock.AssertCalled(t, "DeleteTask", 1)
}

func TestMarkInProgress(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("UpdateTask", mock.AnythingOfType("int"), mock.AnythingOfType("Task")).Return(nil)
	task := models.Task{
		Status: models.Inprogress,
	}

	// Act
	markInProgress(taskServiceMock, []string{"mark-in-progress", "1"})

	// Assert
	taskServiceMock.AssertCalled(t, "UpdateTask", 1, task)
}

func TestMarkDone(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("UpdateTask", mock.AnythingOfType("int"), mock.AnythingOfType("Task")).Return(nil)
	task := models.Task{
		Status: models.Done,
	}

	// Act
	markDone(taskServiceMock, []string{"mark-done", "1"})

	// Assert
	taskServiceMock.AssertCalled(t, "UpdateTask", 1, task)
}

func TestList(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("GetAllTask", mock.Anything).Return([]models.Task{
		{Id: 1, Name: "testing todo", Status: models.Todo},
		{Id: 2, Name: "testing in progress", Status: models.Inprogress},
		{Id: 3, Name: "testing done", Status: models.Done},
	})

	// Act
	list(taskServiceMock, []string{"list"})

	// Assert
	taskServiceMock.AssertCalled(t, "GetAllTask", mock.Anything)
}

func TestListWithStatus(t *testing.T) {
	// Arrange
	taskServiceMock := new(service.TaskServiceMock)
	taskServiceMock.On("GetAllTask", mock.Anything).Return([]models.Task{
		{Id: 1, Name: "testing todo", Status: models.Todo},
		{Id: 2, Name: "testing in progress", Status: models.Inprogress},
		{Id: 3, Name: "testing done", Status: models.Done},
	})
	status := models.Todo

	// Act
	list(taskServiceMock, []string{"list", "todo"})

	// Assert
	taskServiceMock.AssertCalled(t, "GetAllTask", &status)
}

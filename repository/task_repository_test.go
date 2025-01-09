package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"task-tracker-cli/models"
	"task-tracker-cli/storage"
	"testing"
)

func TestTaskRepositoryImpl_Create(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Return(nil)
	fileStorageMock.On("Save", mock.Anything).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}
	task := models.Task{
		Name:   "testing",
		Status: models.Todo,
	}

	// Act
	err := taskRepository.Create(task)

	// Assert
	assert.NoError(t, err)
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
	fileStorageMock.AssertNumberOfCalls(t, "Save", 1)
	fileStorageMock.AssertCalled(t, "Save", []models.Task{
		{Id: 1, Name: "testing", Status: models.Todo},
	})
}

func TestTaskRepositoryImpl_Update(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
		}
	}).Return(nil)
	fileStorageMock.On("Save", mock.Anything).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}
	task := models.Task{
		Name:   "testing",
		Status: models.Todo,
	}

	// Act
	err := taskRepository.Update(1, task)

	// Assert
	assert.NoError(t, err)
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
	fileStorageMock.AssertNumberOfCalls(t, "Save", 1)
	fileStorageMock.AssertCalled(t, "Save", []models.Task{
		{Id: 1, Name: "testing", Status: models.Todo},
	})
}

func TestTaskRepositoryImpl_Delete(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
		}
	}).Return(nil)
	fileStorageMock.On("Save", mock.Anything).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	err := taskRepository.Delete(1)

	// Assert
	assert.NoError(t, err)
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
	fileStorageMock.AssertNumberOfCalls(t, "Save", 1)
	fileStorageMock.AssertCalled(t, "Save", []models.Task{})
}

func TestTaskRepositoryImpl_Get(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
		}
	}).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	task, err := taskRepository.Get(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, task, models.Task{Id: 1, Name: "Task 1"})
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
}

func TestTaskRepositoryImpl_GetAll(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
			{Id: 2, Name: "Task 2", Status: models.Inprogress},
		}
	}).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	tasks := taskRepository.GetAll(nil)

	// Assert
	assert.Equal(t, tasks, []models.Task{
		{Id: 1, Name: "Task 1", Status: models.Todo},
		{Id: 2, Name: "Task 2", Status: models.Inprogress},
	})
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
}

func TestTaskRepositoryImpl_GetAll_Todo(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
			{Id: 2, Name: "Task 2", Status: models.Inprogress},
			{Id: 3, Name: "Task 3", Status: models.Done},
		}
	}).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	modelStatus := models.Todo
	tasks := taskRepository.GetAll(&modelStatus)

	// Assert
	assert.Equal(t, tasks, []models.Task{
		{Id: 1, Name: "Task 1", Status: models.Todo},
	})
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
}

func TestTaskRepositoryImpl_GetAll_InProgress(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
			{Id: 2, Name: "Task 2", Status: models.Inprogress},
			{Id: 3, Name: "Task 3", Status: models.Done},
		}
	}).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	modelStatus := models.Inprogress
	tasks := taskRepository.GetAll(&modelStatus)

	// Assert
	assert.Equal(t, tasks, []models.Task{
		{Id: 2, Name: "Task 2", Status: models.Inprogress},
	})
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
}

func TestTaskRepositoryImpl_GetAll_Done(t *testing.T) {
	// Arrange
	fileStorageMock := new(storage.FileStorageMock)
	fileStorageMock.On("Load", mock.Anything).Run(func(args mock.Arguments) {
		result := args.Get(0).(*[]models.Task)
		*result = []models.Task{
			{Id: 1, Name: "Task 1", Status: models.Todo},
			{Id: 2, Name: "Task 2", Status: models.Inprogress},
			{Id: 3, Name: "Task 3", Status: models.Done},
		}
	}).Return(nil)
	taskRepository := TaskRepositoryImpl{
		Filestorage: fileStorageMock,
	}

	// Act
	modelStatus := models.Done
	tasks := taskRepository.GetAll(&modelStatus)

	// Assert
	assert.Equal(t, tasks, []models.Task{
		{Id: 3, Name: "Task 3", Status: models.Done},
	})
	fileStorageMock.AssertNumberOfCalls(t, "Load", 1)
}

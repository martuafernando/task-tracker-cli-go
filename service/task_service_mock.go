package service

import (
	"github.com/stretchr/testify/mock"
	"task-tracker-cli/models"
)

type TaskServiceMock struct {
	mock.Mock
}

func (s *TaskServiceMock) AddTask(taskname string) error {
	args := s.Called(taskname)
	return args.Error(0)
}

func (s *TaskServiceMock) UpdateTask(id int, task models.Task) error {
	args := s.Called(id, task)
	return args.Error(0)
}

func (s *TaskServiceMock) DeleteTask(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *TaskServiceMock) GetAllTask(status *models.Status) []models.Task {
	args := s.Called(status)
	return args.Get(0).([]models.Task)
}

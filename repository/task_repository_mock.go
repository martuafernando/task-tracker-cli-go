package repository

import (
	"github.com/stretchr/testify/mock"
	"task-tracker-cli/models"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (s *TaskRepositoryMock) Create(task models.Task) error {
	args := s.Called(task)
	return args.Error(0)
}

func (s *TaskRepositoryMock) Update(id int, task models.Task) error {
	args := s.Called(id, task)
	return args.Error(0)
}

func (s *TaskRepositoryMock) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *TaskRepositoryMock) Get(id int) (models.Task, error) {
	args := s.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (s *TaskRepositoryMock) GetAll(status *models.Status) []models.Task {
	args := s.Called(status)
	return args.Get(0).([]models.Task)
}

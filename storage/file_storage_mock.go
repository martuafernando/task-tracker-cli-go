package storage

import (
	"github.com/stretchr/testify/mock"
)

type FileStorageMock struct {
	mock.Mock
}

func (s *FileStorageMock) Save(data any) error {
	args := s.Called(data)
	return args.Error(0)
}

func (s *FileStorageMock) Load(v any) error {
	args := s.Called(v)
	return args.Error(0)
}

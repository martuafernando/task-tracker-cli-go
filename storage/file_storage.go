package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FileStorage interface {
	Save(data any) error
	Load(v any) error
}

type FileStorageImpl struct {
	Filename string
}

func (s *FileStorageImpl) Save(data any) error {
	dir := filepath.Dir(s.Filename)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(s.Filename)

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(data)
}

func (s *FileStorageImpl) Load(v any) error {
	file, err := os.Open(s.Filename)

	if os.IsNotExist(err) {
		emptyData := []byte("[]")
		decoder := json.NewDecoder(bytes.NewReader(emptyData))
		return decoder.Decode(v)
	}

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}

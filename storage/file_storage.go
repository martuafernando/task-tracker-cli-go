package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveToFile(filename string, data any) error {
	dir := filepath.Dir(filename)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(data)
}

func LoadFromFile(filename string, v any) error {
	file, err := os.Open(filename)

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

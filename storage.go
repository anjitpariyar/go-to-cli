package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

// create NewStorage
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// create save method
func (s *Storage[T]) Save(data T) error {
	// convert data to json
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

// read file
func (s *Storage[T]) Read(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}

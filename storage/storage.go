package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
)

type Storage struct {
	path string
}

func New(path string) *Storage {
	return &Storage{path: path}
}

func (s *Storage) Load() ([]model.Task, error) {
	data, err := os.ReadFile(s.path)
	if errors.Is(err, os.ErrNotExist) {
		return []model.Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Storage) Save(tasks []model.Task) error {
	out, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, out, 0644)

}

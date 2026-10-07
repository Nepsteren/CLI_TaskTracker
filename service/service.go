package service

import (
	"errors"
	"time"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
	"github.com/Nepsteren/CLI_TaskTracker.git/storage"
)

var (
	ErrEmptyDesc     = errors.New("описание не может быть пустым")
	ErrTaskNotFound  = errors.New("задача не найдена")
	ErrInvalidStatus = errors.New("неверный статус")
)

type Service struct {
	store *storage.Storage
	tasks []model.Task
}

func (s *Service) nextID() int {
	max := 0
	for _, t := range s.tasks {
		if t.Id > max {
			max = t.Id
		}
	}
	return max + 1
}

func (s *Service) findIndex(id int) int {
	for i, t := range s.tasks {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func New(store *storage.Storage) (*Service, error) {
	tasks, err := store.Load()
	if err != nil {
		return nil, err
	}
	return &Service{store: store, tasks: tasks}, nil
}

func (s *Service) All() []model.Task {
	return s.tasks
}

func (s *Service) ByStatus(status model.Status) []model.Task {
	var res []model.Task
	for _, t := range s.tasks {
		if t.Status == status {
			res = append(res, t)
		}
	}
	return res
}

func (s *Service) Add(desc string) error {
	if desc == "" {
		return ErrEmptyDesc
	}
	now := time.Now()
	s.tasks = append(s.tasks, model.Task{
		Id:          s.nextID(),
		Description: desc,
		Status:      model.StatusUndone,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	return s.store.Save(s.tasks)
}

func (s *Service) Delete(id int) error {

	idx := s.findIndex(id)
	if idx == -1 {
		return ErrTaskNotFound
	}
	s.tasks = append(s.tasks[:idx], s.tasks[idx+1:]...)

	return s.store.Save(s.tasks)
}

func (s *Service) Update(id int, desc string) error {
	if desc == "" {
		return ErrEmptyDesc
	}

	idx := s.findIndex(id)
	if idx == -1 {
		return ErrTaskNotFound
	}

	s.tasks[idx].Description = desc
	s.tasks[idx].UpdatedAt = time.Now()
	return s.store.Save(s.tasks)
}

func (s *Service) Mark(id int, status model.Status) error {
	if !status.Valid() {
		return ErrInvalidStatus
	}

	idx := s.findIndex(id)
	if idx == -1 {
		return ErrTaskNotFound
	}

	s.tasks[idx].Status = status
	s.tasks[idx].UpdatedAt = time.Now()
	return s.store.Save(s.tasks)
}

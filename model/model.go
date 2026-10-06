package model

import "time"

type Status string

const (
	StatusUndone     Status = "undone"
	StatusDone       Status = "done"
	StatusInProgress Status = "in progress"
)

func (s Status) Valid() bool {
	switch s {
	case StatusDone, StatusUndone, StatusInProgress:
		return true
	}
	return false
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

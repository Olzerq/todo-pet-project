package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	return &Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (t *Task) Complete() {
	t.UpdatedAt = time.Now()
	t.Completed = true
}
func (t *Task) Uncomplete() {
	t.UpdatedAt = time.Now()
	t.Completed = false
}
func (t *Task) UpdateTitle(title string) error {
	if title == "" {
		return errors.New("task title cannot be empty")
	}
	t.UpdatedAt = time.Now()
	t.Title = title
	return nil
}
func (t *Task) UpdateDescription(description string) {
	t.UpdatedAt = time.Now()
	t.Description = description
}
func generateID() string {
	return uuid.New().String()
}

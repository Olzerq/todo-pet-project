package repository

import (
	"context"

	"github.com/olzerq/todo-pet-project/internal/domain/entity"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entity.Task) error
	GetByID(ctx context.Context, id string) (*entity.Task, error)
	GetAll(ctx context.Context) ([]*entity.Task, error)
	Update(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, id string) error
}

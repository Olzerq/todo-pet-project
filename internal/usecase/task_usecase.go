package usecase

import (
	"context"
	"fmt"
	"github.com/olzerq/todo-pet-project/internal/domain/entity"
	"github.com/olzerq/todo-pet-project/internal/domain/repository"
)

type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

type UpdateOption func(task *entity.Task) error

func WithTitle(title string) UpdateOption {
	return func(task *entity.Task) error {
		return task.UpdateTitle(title)
	}
}
func WithDescription(description string) UpdateOption {
	return func(task *entity.Task) error {
		task.UpdateDescription(description)
		return nil
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, title, description string) (*entity.Task, error) {
	task, err := entity.NewTask(title, description)
	if err != nil {
		return nil, err
	}

	if err := uc.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}
func (uc *TaskUseCase) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	return uc.taskRepo.GetByID(ctx, id)
}
func (uc *TaskUseCase) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	return uc.taskRepo.GetAll(ctx)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, id string, opts ...UpdateOption) error {
	task, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	for _, opt := range opts {
		if err := opt(task); err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}
	}
	return uc.taskRepo.Update(ctx, task)
}

func (uc *TaskUseCase) CompleteTask(ctx context.Context, id string) error {
	task, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	task.Complete()

	return uc.taskRepo.Update(ctx, task)
}

func (uc *TaskUseCase) UncompleteTask(ctx context.Context, id string) error {
	task, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	task.Uncomplete()

	return uc.taskRepo.Update(ctx, task)
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	return uc.taskRepo.Delete(ctx, id)
}

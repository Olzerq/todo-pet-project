package postgres

import (
	"context"
	"database/sql"
	"github.com/olzerq/todo-pet-project/internal/domain/entity"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *entity.Task) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO tasks (id, title, description, completed, created_at, updated_at) VALUES (:id, :title, :description, :completed, :created_at, :updated_at)`,
		sql.Named("id", task.ID),
		sql.Named("title", task.Title),
		sql.Named("description", task.Description),
		sql.Named("completed", task.Completed),
		sql.Named("created_at", task.CreatedAt),
		sql.Named("updated_at", task.UpdatedAt))
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*entity.Task, error) {
	t := &entity.Task{}
	err := r.db.QueryRowContext(ctx, `SELECT * FROM tasks WHERE id = :id`, sql.Named("id", id)).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Completed,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*entity.Task, error) {
	var tasks []*entity.Task
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t *entity.Task
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *entity.Task) error {
	_, err := r.db.ExecContext(ctx, `UPDATE tasks SET title = :title, description = :description, completed = :completed, updated_at = :updated_at WHERE id = :id`,
		sql.Named("id", task.ID),
		sql.Named("title", task.Title),
		sql.Named("description", task.Description),
		sql.Named("completed", task.Completed),
		sql.Named("updated_at", task.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM tasks WHERE id = :id`, sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}

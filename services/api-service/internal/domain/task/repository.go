package task

import (
	"context"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(context.Context, *Task) error

	GetManyByUserID(context.Context, uuid.UUID) (*[]Task, error)
	GetByID(context.Context, uuid.UUID) (*Task, error)
}

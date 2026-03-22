package task

import (
	"api-service/internal/domain/shared"
	domainTask "api-service/internal/domain/task"
	domainUser "api-service/internal/domain/user"
	"context"

	"github.com/google/uuid"
)

type create_UseCase struct {
	userRepo domainUser.UserRepository
	taskRepo domainTask.TaskRepository
}

// return created task id, or error
func (uc *create_UseCase) CreateNewTask(ctx context.Context, userID uuid.UUID, _type string, payload []byte) (string, error) {
	u, err := uc.userRepo.GetUserByID(userID)
	if err != nil {
		return "", err
	}

	t, err := u.CreateTask(domainUser.CreateTaskCmd{
		Type:    shared.TaskType(_type),
		Payload: payload,
	})
	if err != nil {
		return "", err
	}

	if err := uc.taskRepo.Save(ctx, t); err != nil {
		return "", err
	}

	return t.GetID().String(), nil
}

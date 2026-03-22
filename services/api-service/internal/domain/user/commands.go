package user

import "api-service/internal/domain/shared"

// task creation command
type CreateTaskCmd struct {
	Type    shared.TaskType
	Payload []byte
}

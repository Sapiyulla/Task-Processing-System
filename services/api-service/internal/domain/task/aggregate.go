package task

import (
	"api-service/internal/domain/shared"

	"github.com/google/uuid"
)

type Task struct {
	id      uuid.UUID
	userID  uuid.UUID
	_type   shared.TaskType
	payload []byte
	status  shared.TaskStatus
}

func NewTask(userID uuid.UUID, _type shared.TaskType, payload []byte) (*Task, error) {
	if !validType(_type) {
		return nil, ErrUndefinedTaskType
	}
	if len(payload) > (10 * 1024 * 1024) {
		return nil, ErrPayloadTooLarge
	}
	return &Task{
		id:      uuid.New(),
		userID:  userID,
		_type:   _type,
		payload: payload,
		status:  shared.CREATED,
	}, nil
}

func (u *Task) GetID() uuid.UUID {
	return u.id
}

func (u *Task) GetUserID() uuid.UUID {
	return u.userID
}

func (u *Task) GetType() shared.TaskType {
	return u._type
}

func (u *Task) GetStatus() shared.TaskStatus {
	return u.status
}

func (u *Task) Run() error {
	switch u.status {
	case shared.FAILED, shared.CREATED:
		u.status = shared.WAIT
		return nil
	default:
		return ErrTaskNotExecutable
	}
}

func validType(t shared.TaskType) bool {
	switch t {
	case shared.AI_Pictire_Prepare:
		return true
	case shared.Data_Processing:
		return true
	case shared.Report_Generate:
		return true
	default:
		return false
	}
}

func validStatus(s shared.TaskStatus) bool {
	switch s {
	case shared.CREATED:
		return true
	case shared.WAIT:
		return true
	case shared.QUEUED:
		return true
	case shared.IN_PROGRESS:
		return true
	case shared.COMPLETED:
		return true
	case shared.FAILED:
		return true
	default:
		return false
	}
}

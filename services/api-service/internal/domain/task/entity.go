package task

import (
	"github.com/google/uuid"
)

type TaskType string

const (
	AI_Pictire_Prepare TaskType = "ai_picture_prepare"
	Data_Processing    TaskType = "data_processing"
	Report_Generate    TaskType = "report_generate"
)

type TaskStatus string

const (
	CREATED     TaskStatus = "CREATED"
	WAIT        TaskStatus = "WAIT"
	QUEUED      TaskStatus = "QUEUED"
	IN_PROGRESS TaskStatus = "IN PROGRESS"
	COMPLETED   TaskStatus = "COMPLETED"
)

type Task struct {
	id      uuid.UUID
	_type   TaskType
	payload []byte
	status  TaskStatus
}

func NewTask(Type TaskType, payload []byte, status TaskStatus) (*Task, error) {
	if len(payload) > 10*1024*1024 {
		return nil, HeavyPayloadErr
	}
	if !validType(Type) {
		return nil, UndefinedTaskType
	}
	if !validStatus(status) {
		return nil, UndefinedTaskStatus
	}
	return &Task{
		id:    uuid.New(),
		_type: Type,
	}, nil
}

func validType(t TaskType) bool {
	switch t {
	case AI_Pictire_Prepare:
		return true
	case Data_Processing:
		return true
	case Report_Generate:
		return true
	default:
		return false
	}
}

func validStatus(s TaskStatus) bool {
	switch s {
	case CREATED:
		return true
	case WAIT:
		return true
	case QUEUED:
		return true
	case IN_PROGRESS:
		return true
	case COMPLETED:
		return true
	default:
		return false
	}
}

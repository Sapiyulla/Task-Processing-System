package task

import "errors"

var (
	ErrPayloadTooLarge     error = errors.New("payload is too large")
	ErrUndefinedTaskType   error = errors.New("task type is undefined")
	ErrUndefinedTaskStatus error = errors.New("task status is undefined")
	ErrTaskNotExecutable   error = errors.New("task not executable")
)

var (
	ErrTaskNotFound error = errors.New("task not found")
)

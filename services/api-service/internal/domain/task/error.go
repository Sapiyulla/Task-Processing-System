package task

import "errors"

var (
	HeavyPayloadErr     error = errors.New("payload is heavy")
	UndefinedTaskType   error = errors.New("task type is undefined")
	UndefinedTaskStatus error = errors.New("task status is undefined")
)

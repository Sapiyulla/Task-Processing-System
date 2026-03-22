package shared

type TaskStatus string

const (
	CREATED     TaskStatus = "CREATED"
	WAIT        TaskStatus = "WAIT"
	QUEUED      TaskStatus = "QUEUED"
	IN_PROGRESS TaskStatus = "IN_PROGRESS"
	COMPLETED   TaskStatus = "COMPLETED"
	FAILED      TaskStatus = "FAILED"
)

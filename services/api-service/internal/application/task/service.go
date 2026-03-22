package task

import domainTask "api-service/internal/domain/task"

type TaskService struct {
	repo domainTask.TaskRepository
}

func NewTaskService(repo domainTask.TaskRepository) *TaskService {
	return &TaskService{}
}

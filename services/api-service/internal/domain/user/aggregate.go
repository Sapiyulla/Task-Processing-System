package user

import (
	"api-service/internal/domain/task"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	id                                           uuid.UUID
	name                                         string
	email                                        string
	tasksCount, dailyLimit, inProgressTasksCount uint16
}

func NewUser(name, email string) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	return &User{
		id:    uuid.New(),
		name:  name,
		email: email,
	}, nil
}

func Rehydrate(id uuid.UUID, name, email string, tasksCount, dailyLimit, inProgressTasks uint16) *User {
	return &User{
		id:                   id,
		name:                 name,
		email:                email,
		tasksCount:           tasksCount,
		dailyLimit:           dailyLimit,
		inProgressTasksCount: inProgressTasks,
	}
}

func (u *User) canCreateTask() error {
	if u.tasksCount >= 70 {
		return ErrTasksLimit
	}
	return nil
}

func (u *User) canStartTask() error {
	if u.dailyLimit >= 10 {
		return ErrDailyLimit
	}
	if u.inProgressTasksCount >= 3 {
		return ErrRunningTasksLimit
	}
	return nil
}

func (u *User) CreateTask(cmd CreateTaskCmd) (*task.Task, error) {
	if err := u.canCreateTask(); err != nil {
		return nil, err
	}
	t, err := task.NewTask(u.id, cmd.Type, cmd.Payload)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func validateName(name string) error {
	if len(name) < 3 || len(name) > 12 {
		return ErrNameLength
	}
	if !regexp.MustCompile(`^[A-Za-z ]+$`).MatchString(name) {
		return ErrNameChars
	}
	if len(strings.Fields(name)) > 3 {
		return ErrNameWordCount
	}
	return nil
}

func validateEmail(email string) error {
	if !regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`).MatchString(email) {
		return ErrEmailInvalid
	}
	return nil
}

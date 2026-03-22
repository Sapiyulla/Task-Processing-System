package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(context.Context, *User) error

	GetUserByID(uuid.UUID) (*User, error)
	GetUserByEmail(string) (*User, error)
}

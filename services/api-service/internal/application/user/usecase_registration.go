package user

import (
	domainUser "api-service/internal/domain/user"
	"context"
)

type registration_UseCase struct {
	repo domainUser.UserRepository
}

// return token (as string), or error
func (uc *registration_UseCase) Registration(ctx context.Context, name, email string) (string, error) {
	if u, err := uc.repo.GetUserByEmail(email); err == nil && u != nil {
		return "", domainUser.ErrUserNotFound
	}
	if ctx.Err() != nil {
		return "", context.Canceled
	}

	u, err := domainUser.NewUser(name, email)
	if err != nil {
		return "", err
	}

	if err := uc.repo.Save(ctx, u); err != nil {
		return "", err
	}
	if ctx.Err() != nil {
		return "", context.Canceled
	}

	// token generation process...

	return "", nil
}

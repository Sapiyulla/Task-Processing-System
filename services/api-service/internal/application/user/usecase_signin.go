package user

import (
	domainUser "api-service/internal/domain/user"
	"context"
)

type signIn_UseCase struct {
	repo domainUser.UserRepository
}

// return token (as string), or error
func (uc *signIn_UseCase) SignIn(ctx context.Context, email string) (string, error) {
	u, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// token generation process...
	_ = u

	return "", nil
}

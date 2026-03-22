package user

import domainUser "api-service/internal/domain/user"

type UserService struct {
	// build with use all use cases
	*registration_UseCase
	*signIn_UseCase
}

func NewUserService(repo domainUser.UserRepository) *UserService {
	return &UserService{
		registration_UseCase: &registration_UseCase{
			repo: repo,
		},
		signIn_UseCase: &signIn_UseCase{
			repo: repo,
		},
	}
}

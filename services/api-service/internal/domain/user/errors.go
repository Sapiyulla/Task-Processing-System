package user

import (
	"errors"
	"net/http"
)

// # domain error
var (
	ErrDailyLimit        error = errors.New("daily tasks limit reached")
	ErrTasksLimit        error = errors.New("task limit has been reached")
	ErrRunningTasksLimit error = errors.New("maximum number of running tasks has now been reached")

	ErrNameLength    error = errors.New("name length must be between 3 and 12 characters")
	ErrNameChars     error = errors.New("name must contain only Latin letters and spaces")
	ErrNameWordCount error = errors.New("name must have at most 3 words")
	ErrEmailInvalid  error = errors.New("invalid email format")
)

// # repository error
var (
	ErrUserAlreadyExists error = errors.New("user already exists")
	ErrUserNotFound      error = errors.New("user not found")
)

func HTTPStatusCode(e error) int {
	switch e {
	case ErrDailyLimit,
		ErrRunningTasksLimit,
		ErrTasksLimit,
		ErrNameWordCount,
		ErrNameChars,
		ErrNameLength,
		ErrEmailInvalid:
		return http.StatusBadRequest
	case ErrUserNotFound:
		return http.StatusNotFound
	case ErrUserAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

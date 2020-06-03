package user

import (
	"github.com/pkg/errors"
)

var (
	// ErrAuthenticationFailure auth failure
	ErrAuthenticationFailure = errors.New("Authentication failed")
	ErrorNotFound            = errors.New("Entity not found")
	ErrForbidden             = errors.New("Attempted action is not allowed")

	// ErrResetExpired occurs when the reset hash exceeds the expiration
	ErrResetExpired = errors.New("Reset expired")
)

// GetUserFromCreateRequest returns user object from request
func (user *UserCreateRequest) GetUserFromCreateRequest() (*User, error) {
	if user == nil {
		return nil, nil
	}

	usr := &User{
		Name: user.Name,
	}

	return usr, nil
}

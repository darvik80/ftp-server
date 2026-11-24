package engine

import "errors"

var (
	ErrAuthFailed         = errors.New("auth failed")           // Authentication failed
	ErrUserNotFound       = errors.New("user not found")        // User not found
	ErrInvalidCredentials = errors.New("invalid credentials")   // Invalid credentials
	ErrUserAlreadyExists  = errors.New("user already exists")   // User already exists
	ErrInternalServer     = errors.New("internal server error") // Internal server error
)

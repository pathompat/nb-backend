package helper

import (
	"errors"
)

var (
	ErrUnauthorized    = errors.New("Unauthorized")
	ErrDatabaseProcess = errors.New("Database error")

	ErrInvalidUserOrPwd = errors.New("Invalid username or password")
)

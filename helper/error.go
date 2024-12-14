package helper

import (
	"errors"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrSelectRecord = errors.New("Database: select error")
	ErrInsertRecord = errors.New("Database: insert error")

	ErrInvalidUserOrPwd = errors.New("Invalid username or password")
	ErrHashPassword     = errors.New("Error hashing password")
	ErrInvalidToken     = errors.New("Invalid token")
	ErrMissingToken     = errors.New("Missing authorization")
	ErrInvalidPathParam = errors.New("Invalid path param")

	ErrForbidden = errors.New("Forbidden")
)

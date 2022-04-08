package errors

import "errors"

var (
	ErrUserCheckFailed       = errors.New("user check failed")
	ErrCommandNotFound       = errors.New("command not found")
	ErrCommandNoErrorHandler = errors.New("command has no error handler")

	ErrUserBanned          = errors.New("user is banned")
	ErrCommandIsGuildOnly  = errors.New("command is guild only")
	ErrCommandIsDirectOnly = errors.New("command is direct only")
	ErrCommandIsAdminOnly  = errors.New("command is admin only")
)

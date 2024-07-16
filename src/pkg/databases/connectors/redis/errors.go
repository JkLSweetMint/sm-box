package redis

import "errors"

var (
	ErrDatabaseNameIsNotSpecified = errors.New("The database name is not specified. ")
	ErrUserIsNotSpecified         = errors.New("The user is not specified. ")
	ErrUserPasswordIsNotSpecified = errors.New("The user password is not specified. ")
	ErrHostIsNotSpecified         = errors.New("The host is not specified. ")
	ErrPortIsNotSpecified         = errors.New("The port is not specified. ")
)

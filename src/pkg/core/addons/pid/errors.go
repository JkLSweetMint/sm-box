package pid

import "errors"

var (
	ErrPidFileAlreadyExist = errors.New("PID file already exist. ")
	ErrPidFileNotExist     = errors.New("PID file not exist. ")
)

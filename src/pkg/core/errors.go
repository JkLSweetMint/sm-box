package core

import "errors"

var (
	ErrInvalidSystemCoreState                    = errors.New("Invalid state of the system core. ")
	ErrSystemCoreIsNotServe                      = errors.New("The system core is not serve. ")
	ErrSystemCoreIsNotBooted                     = errors.New("The system core is not booted. ")
	ErrSystemCoreIsAlreadyBooted                 = errors.New("The system core is already booted. ")
	ErrISystemCoresAlreadyTurnedOff              = errors.New("The system core is already turned off. ")
	ErrSystemCoreIsAlreadyServed                 = errors.New("The system core is already served. ")
	ErrSystemCoreAlreadyClosedStartIsNotPossible = errors.New("The system core has already been stopped the launch is not possible. ")
)

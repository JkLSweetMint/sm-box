package error_list

import (
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// I-190001
var (
	UserCouldNotBeAuthorizedOnRemoteService = c_errors.Constructor[c_errors.Error]{
		ID:     "I-190001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user could not be authorized on the remote service "),
	}.Build()
)

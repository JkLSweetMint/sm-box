package service

import (
	"sm-box/internal/services/user_manager/transports/rest_api"
)

// Transports - описание транспортной части сервиса.
type Transports interface {
	RestApi() rest_api.Engine
}

// components - транспортная часть сервиса.
type transports struct {
	restApi rest_api.Engine
}

// RestApi - получение http rest api.
func (t *transports) RestApi() rest_api.Engine {
	return t.restApi
}

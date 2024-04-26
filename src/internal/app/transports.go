package app

import (
	"sm-box/internal/app/transports/rest_api"
)

// Transports - описание транспортной части коробки.
type Transports interface {
	RestApi() rest_api.Engine
}

// components - транспортная часть коробки.
type transports struct {
	restApi rest_api.Engine
}

// RestApi - получение http rest api.
func (t *transports) RestApi() rest_api.Engine {
	return t.restApi
}

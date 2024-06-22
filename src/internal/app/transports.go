package app

import (
	"sm-box/internal/app/transports/http_proxy"
)

// Transports - описание транспортной части коробки.
type Transports interface {
	HttpProxy() http_proxy.Engine
}

// components - транспортная часть коробки.
type transports struct {
	httpProxy http_proxy.Engine
}

// HttpProxy - получение http http proxy.
func (t *transports) HttpProxy() http_proxy.Engine {
	return t.httpProxy
}

package service

import (
	"sm-box/internal/services/i18n/transport/servers/http_rest_api"
)

// Transport - описание транспортной части сервиса.
type Transport interface {
	Servers() TransportServers
	Gateways() TransportGateways
}

// TransportServers - описание серверов транспортной части сервиса.
type TransportServers interface {
	HttpRestApi() http_rest_api.Server
}

// TransportGateways - описание шлюзов транспортной части сервиса.
type TransportGateways interface{}

// components - транспортная часть сервиса.
type transport struct {
	servers  *transportServers
	gateways *transportGateways
}

// transportsServers - сервера транспортной части сервиса.
type transportServers struct {
	httpRestApi http_rest_api.Server
}

// transportsGateways - шлюзы транспортной части сервиса.
type transportGateways struct {
}

// Servers - получение серверов транспортной части сервиса.
func (t *transport) Servers() TransportServers {
	return t.servers
}

// Gateways - получение шлюзов транспортной части сервиса.
func (t *transport) Gateways() TransportGateways {
	return t.gateways
}

// HttpRestApi - получение http rest api сервера.
func (t *transportServers) HttpRestApi() http_rest_api.Server {
	return t.httpRestApi
}

package service

import (
	"sm-box/internal/services/authentication/transport/servers/grpc_authentication_srv"
	"sm-box/internal/services/authentication/transport/servers/http_rest_api"
)

// Transport - описание транспортной части сервиса.
type Transport interface {
	Servers() TransportServers
	Gateways() TransportGateways
}

// TransportServers - описание серверов транспортной части сервиса.
type TransportServers interface {
	HttpRestApi() http_rest_api.Server
	GrpcAuthenticationService() grpc_authentication_srv.Server
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
	httpRestApi               http_rest_api.Server
	grpcAuthenticationService grpc_authentication_srv.Server
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

// GrpcAuthenticationService - получение сервера grpc сервиса аутентификации.
func (t *transportServers) GrpcAuthenticationService() grpc_authentication_srv.Server {
	return t.grpcAuthenticationService
}

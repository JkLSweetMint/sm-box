package service

import (
	basic_authentication_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/basic_authentication_service"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
	users_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/users_service"
	http_rest_api "sm-box/internal/services/authentication/transport/servers/http/rest_api"
)

// Transport - описание транспортной части сервиса.
type Transport interface {
	Servers() TransportServers
	Gateways() TransportGateways
}

// TransportServers - описание серверов транспортной части сервиса.
type TransportServers interface {
	Http() TransportServersHttp
}

// TransportServersHttp - описание серверов транспортной части сервиса по http.
type TransportServersHttp interface {
	RestApi() http_rest_api.Server
}

// TransportGateways - описание шлюзов транспортной части сервиса.
type TransportGateways interface {
	BasicAuthenticationService() *basic_authentication_service_gateway.Gateway
	ProjectService() *projects_service_gateway.Gateway
	UsersService() *users_service_gateway.Gateway
}

// --------- internal ---------

// transport - транспортная часть сервиса.
type transport struct {
	servers  *transportServers
	gateways *transportGateways
}

// transportsServers - сервера транспортной части сервиса.
type transportServers struct {
	http *transportServersHttp
}

// transportsServers - сервера транспортной части сервиса по http.
type transportServersHttp struct {
	restApi http_rest_api.Server
}

// transportsGateways - шлюзы транспортной части сервиса.
type transportGateways struct {
	basicAuthenticationService *basic_authentication_service_gateway.Gateway
	projectService             *projects_service_gateway.Gateway
	usersService               *users_service_gateway.Gateway
}

// Servers - получение серверов транспортной части сервиса.
func (t *transport) Servers() TransportServers {
	return t.servers
}

// Gateways - получение шлюзов транспортной части сервиса.
func (t *transport) Gateways() TransportGateways {
	return t.gateways
}

// Http - получение серверов транспортной части сервиса по http.
func (t *transportServers) Http() TransportServersHttp {
	return t.http
}

// RestApi - получение http rest api сервера.
func (t *transportServersHttp) RestApi() http_rest_api.Server {
	return t.restApi
}

// BasicAuthenticationService - получение шлюза транспортной части сервиса.
func (t *transportGateways) BasicAuthenticationService() *basic_authentication_service_gateway.Gateway {
	return t.basicAuthenticationService
}

// ProjectService - получение шлюза транспортной части сервиса.
func (t *transportGateways) ProjectService() *projects_service_gateway.Gateway {
	return t.projectService
}

// UsersService - получение шлюза транспортной части сервиса.
func (t *transportGateways) UsersService() *users_service_gateway.Gateway {
	return t.usersService
}

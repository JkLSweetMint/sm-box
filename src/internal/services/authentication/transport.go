package service

import (
	authentication_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/authentication_service"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
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
	AuthenticationService() *authentication_service_gateway.Gateway
	ProjectService() *projects_service_gateway.Gateway
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
	authenticationService *authentication_service_gateway.Gateway
	projectService        *projects_service_gateway.Gateway
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

// AuthenticationService - получение шлюза транспортной части сервиса.
func (t *transportGateways) AuthenticationService() *authentication_service_gateway.Gateway {
	return t.authenticationService
}

// ProjectService - получение шлюза транспортной части сервиса.
func (t *transportGateways) ProjectService() *projects_service_gateway.Gateway {
	return t.projectService
}

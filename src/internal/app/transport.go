package app

import (
	grpc_authentication_srv "sm-box/internal/app/transport/servers/grpc/authentication_service"
	grpc_projects_srv "sm-box/internal/app/transport/servers/grpc/projects_service"
	http_rest_api "sm-box/internal/app/transport/servers/http/rest_api"
)

// Transport - описание транспортной части приложения.
type Transport interface {
	Servers() TransportServers
	Gateways() TransportGateways
}

// TransportServers - описание серверов транспортной части приложения.
type TransportServers interface {
	Http() TransportServersHttp
	Grpc() TransportServersGrpc
}

// TransportServersHttp - описание серверов транспортной части приложения по http.
type TransportServersHttp interface {
	RestApi() http_rest_api.Server
}

// TransportServersGrpc - описание серверов транспортной части приложения по grpc.
type TransportServersGrpc interface {
	ProjectsService() grpc_projects_srv.Server
	AuthenticationService() grpc_authentication_srv.Server
}

// TransportGateways - описание шлюзов транспортной части приложения.
type TransportGateways interface{}

// --------- internal ---------

// transport - транспортная часть приложения.
type transport struct {
	servers  *transportServers
	gateways *transportGateways
}

// transportsServers - сервера транспортной части приложения.
type transportServers struct {
	http *transportServersHttp
	grpc *transportServersGrpc
}

// transportServersHttp - сервера транспортной части приложения по http.
type transportServersHttp struct {
	restApi http_rest_api.Server
}

// transportServersGrpc - сервера транспортной части приложения по grpc.
type transportServersGrpc struct {
	projectsService       grpc_projects_srv.Server
	authenticationService grpc_authentication_srv.Server
}

// transportsGateways - шлюзы транспортной части приложения.
type transportGateways struct {
}

// Servers - получение серверов транспортной части приложения.
func (t *transport) Servers() TransportServers {
	return t.servers
}

// Gateways - получение шлюзов транспортной части приложения.
func (t *transport) Gateways() TransportGateways {
	return t.gateways
}

// Http - получение серверов транспортной части приложения по http.
func (t *transportServers) Http() TransportServersHttp {
	return t.http
}

// Grpc - получение серверов транспортной части приложения по grpc.
func (t *transportServers) Grpc() TransportServersGrpc {
	return t.grpc
}

// RestApi - получение http rest api сервера.
func (t *transportServersHttp) RestApi() http_rest_api.Server {
	return t.restApi
}

// ProjectsService - получение сервера для приложения проектов системы.
func (t *transportServersGrpc) ProjectsService() grpc_projects_srv.Server {
	return t.projectsService
}

// AuthenticationService - получение сервера для приложения проектов системы.
func (t *transportServersGrpc) AuthenticationService() grpc_authentication_srv.Server {
	return t.authenticationService
}

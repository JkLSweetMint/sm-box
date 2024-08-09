package service

import (
	popup_notifications_srv "sm-box/internal/services/notifications/transport/servers/grpc/popup_notifications_service"
	user_notifications_srv "sm-box/internal/services/notifications/transport/servers/grpc/user_notifications_service"
	http_rest_api "sm-box/internal/services/notifications/transport/servers/http/rest_api"
)

// Transport - описание транспортной части сервиса.
type Transport interface {
	Servers() TransportServers
	Gateways() TransportGateways
}

// TransportServers - описание серверов транспортной части сервиса.
type TransportServers interface {
	Http() TransportServersHttp
	Grpc() TransportServersGrpc
}

// TransportServersHttp - описание серверов транспортной части сервиса по http.
type TransportServersHttp interface {
	RestApi() http_rest_api.Server
}

// TransportServersGrpc - описание серверов транспортной части приложения по grpc.
type TransportServersGrpc interface {
	UserNotificationsService() user_notifications_srv.Server
	PopupNotificationsService() popup_notifications_srv.Server
}

// TransportGateways - описание шлюзов транспортной части сервиса.
type TransportGateways interface{}

// --------- internal ---------

// transport - транспортная часть сервиса.
type transport struct {
	servers  *transportServers
	gateways *transportGateways
}

// transportsServers - сервера транспортной части сервиса.
type transportServers struct {
	http *transportServersHttp
	grpc *transportServersGrpc
}

// transportsServers - сервера транспортной части сервиса по http.
type transportServersHttp struct {
	restApi http_rest_api.Server
}

// transportServersGrpc - сервера транспортной части приложения по grpc.
type transportServersGrpc struct {
	userNotificationsService  user_notifications_srv.Server
	popupNotificationsService popup_notifications_srv.Server
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

// Http - получение серверов транспортной части сервиса по http.
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

// UserNotificationsService - получение grpc сервера.
func (t *transportServersGrpc) UserNotificationsService() user_notifications_srv.Server {
	return t.userNotificationsService
}

// PopupNotificationsService - получение grpc сервера.
func (t *transportServersGrpc) PopupNotificationsService() popup_notifications_srv.Server {
	return t.popupNotificationsService
}

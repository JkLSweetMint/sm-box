package grpc_authentication_srv

import (
	"context"
	"fmt"
	pb "sm-box/transport/proto/pb/golang/authentication"
)

// RegisterHttpRoutes - регистрация http маршрутов.
func (srv *server) RegisterHttpRoutes(ctx context.Context, request *pb.RegisterHttpRouteRequest) (response *pb.EmptyResponse, err error) {
	response = new(pb.EmptyResponse)

	fmt.Printf("%+v\n", request)

	return
}

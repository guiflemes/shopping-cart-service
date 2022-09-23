package ports

import (
	"context"

	"shopping_cart/src/app"
	"shopping_cart/src/genproto"
)

type grpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) grpcServer {
	return grpcServer{app: application}
}

func (grpc grpcServer) AddItem(ctx context.Context, request *genproto.AddItemRequest) {}

package ports

import (
	"context"
	"shopping_cart/src/app"
	"shopping_cart/src/app/command"
	"shopping_cart/src/genproto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) grpcServer {
	return grpcServer{app: application}
}

func (g grpcServer) AddItem(ctx context.Context, request *genproto.AddItemRequest) (*genproto.Cart, error) {
	addItem := g.app.Commands.AddItem
	cmd := command.AddItem{CartId: request.CartId, ItemId: request.ItemId, Quantity: request.Quantity}

	if err := addItem.Handle(ctx, cmd); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &genproto.Cart{}, nil
}

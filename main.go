package main

import (
	"context"
	"shopping_cart/src/common/server"
	"shopping_cart/src/genproto"
	"shopping_cart/src/ports"
	"shopping_cart/src/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunGrpcServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		genproto.RegisterShoppingCartServiceServer(server, svc)
		reflection.Register(server)
	})

}

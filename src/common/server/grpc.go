package server

import (
	"fmt"
	"net"
	"os"
	"shopping_cart/src/common/logs"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	logger := logrus.New()
	logs.SetFormatter(logger)
	logger.SetLevel(logrus.WarnLevel)

	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logger))

}

func RunGrpcServer(register func(server *grpc.Server)) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	runServerOnAddr(addr, register)

}

func runServerOnAddr(addr string, register func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),
		),
	)

	register(grpcServer)

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithField("grpcEndpoint", addr).Info("Starting: gRPC listener")
	logrus.Fatal(grpcServer.Serve(listen))
}

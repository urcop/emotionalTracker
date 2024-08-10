package app

import (
	"fmt"
	"github.com/FoodMoodOTG/examplearch/services/config"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GRPCServer struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func NewGRPCServer() *GRPCServer {
	domainCtx := InitCtx().Make()
	cfg := config.Make()

	address := fmt.Sprintf("%s:%s", cfg.GrpcHost(), cfg.GrpcPort())

	listener, err := net.Listen("tcp", address)
	if err != nil {
		domainCtx.Services().Logger().Error(
			"error initializing grpc server",
			"op", "grpc_server.NewGRPCServer", "error", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	//sso_v1.RegisterExampleServiceServer(grpcServer, handlers.NewHandler(domainCtx))

	domainCtx.Services().Logger().Info(
		"Starting grpc server",
		"op", "grpc_server.NewGRPCServer", "address", address)
	return &GRPCServer{
		listener:   listener,
		grpcServer: grpcServer,
	}
}

func (s *GRPCServer) Start() {
	err := s.grpcServer.Serve(s.listener)
	if err != nil {
		slog.Error("failed starting GRPC server", "error", err)
	}
}

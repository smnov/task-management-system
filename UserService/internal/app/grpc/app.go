package grpcapp

import "google.golang.org/grpc"

type App struct {
	gRPCServer *grpc.Server
}

func New() *App {
	server := grpc.NewServer()
	return &App{
		gRPCServer: server,
	}
}

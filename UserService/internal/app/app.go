package app

import (
	grpcapp "user_service/internal/app/grpc"

	"github.com/sirupsen/logrus"
)

type Application struct {
	gRPCServer *grpcapp.App
	logger  *logrus.Logger
}

func New(logger *logrus.Logger) *Application {
	app := grpcapp.New()
	return &Application{
		gRPCServer: 	app,
		logger:     logger,
	}
}

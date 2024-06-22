package server

import (
	"context"

	"google.golang.org/grpc"
)

type Auth interface {
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
	Login(ctx context.Context, email string, password string) (token string, err error)
}


type serverAPI struct {
	userv1.UnimplementedAuthServer
	auth Auth
}

func Register(grpcServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(grpcServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	return &ssov1.RegisterResponse{}, nil
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	token, err := s.auth.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	resp := &ssov1.LoginResponse{Token: token}
	return resp, nil
}

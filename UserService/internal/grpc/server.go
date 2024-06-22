package server

import (
	"context"

	"google.golang.org/grpc"
	userv1 "github.com/smnov/task-management-system/Protos/protos/user_service_v1"
)

type Auth interface {
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
	Login(ctx context.Context, email string, password string) (token string, err error)
}


type serverAPI struct {
	userv1.UnimplementedUserV1Server
	auth Auth
}

func Register(grpcServer *grpc.Server, auth Auth) {
	userv1.RegisterUserV1Server(grpcServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Register(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	return &userv1.CreateUserResponse{}, nil
}

func (s *serverAPI) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	token, err := s.auth.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	resp := &userv1.LoginResponse{Token: token}
	return resp, nil
}

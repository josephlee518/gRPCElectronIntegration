package user

import (
	"context"
	pb "github.com/gRPCElectronIntegration/proto"
)

type user struct {
	pb.UnimplementedUserServer
}

// TODO: Login
func (s *user) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {

}

// TODO: Register
func (s *user) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

}

// TODO: GetUserInfo
func (s *user) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {

}

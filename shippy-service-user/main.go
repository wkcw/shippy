package main

import (
	"context"

	pb "github.com/wkcw/shippy/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService()
}

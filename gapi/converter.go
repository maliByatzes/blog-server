package gapi

import (
	db "github.com/maliByatzes/blog-server/db/sqlc"
	"github.com/maliByatzes/blog-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		Email:     user.Email,
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

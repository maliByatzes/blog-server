package gapi

import (
	"fmt"

	db "github.com/maliByatzes/blog-server/db/sqlc"
	"github.com/maliByatzes/blog-server/pb"
	"github.com/maliByatzes/blog-server/token"
	"github.com/maliByatzes/blog-server/util"
)

// Server serves gRPC requests for our blog server service
type Server struct {
	pb.UnimplementedBlogServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server y
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/maliByatzes/blog-server/db/sqlc"
	"github.com/maliByatzes/blog-server/token"
	"github.com/maliByatzes/blog-server/util"
)

// Server serves up HTTP requests
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new server and sets up config
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create a new token: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setUpRoutes()

	return server, nil
}

func (server *Server) setUpRoutes() {
	router := gin.Default()

	// Non-protected routes
	router.POST("/users/register", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.router = router
}

// Start runs the HTTP server on a specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

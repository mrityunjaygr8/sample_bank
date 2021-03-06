package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mrityunjaygr8/sample_bank/db/sqlc"
)

// Server serves HTTP requests for the banking service
type Server struct {
	store  db.Store
	router gin.Engine
}

// NewServer creates a new server instance with the specified store
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", currencyValid)
	}

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccounts)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.POST("transfers", server.createTransfer)

	server.router = *router
	return server
}

// Start starts the http server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func erroresponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

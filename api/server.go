package api

import (
	db "github.com/YuanData/SharedBoard/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.POST("/sharedlink", server.createSharedlink)
	router.GET("/sharedlink/:id", server.getSharedlink)
	router.GET("/sharedlinks", server.listSharedlinks)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

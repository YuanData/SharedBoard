package api

import (
	"strings"

	db "github.com/YuanData/SharedBoard/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	allowedOrigins := []string{
		".github.io",
		// "http://localhost:1313",
	}
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if strings.HasSuffix(origin, allowedOrigin) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Next()
	})

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

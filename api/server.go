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

	router.POST("/sharedlinks", server.createSharedlink) //新增sharedlink
	router.GET("/sharedlinks/:id", server.getSharedlink) //返回一個sharedlink
	router.GET("/sharedlinks", server.listSharedlinks)   //返回 list of sharedlink

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

package api

import (
	"net/http"

	db "github.com/YuanData/SharedBoard/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSharedlinkRequest struct {
	Name    string `json:"name" binding:"required"`
	Urlhash string `json:"urlhash" binding:"required"`
}

func (server *Server) createSharedlink(ctx *gin.Context) {
	var req createSharedlinkRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSharedlinkParams{
		Name:    req.Name,
		Urlhash: req.Urlhash,
	}

	sharedlink, err := server.store.CreateSharedlink(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sharedlink)
}

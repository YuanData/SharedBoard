package api

import (
	"database/sql"
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

type getSharedlinkRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSharedlink(ctx *gin.Context) {
	var req getSharedlinkRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sharedlink, err := server.store.GetSharedlink(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sharedlink)
}

type listSharedlinkRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSharedlinks(ctx *gin.Context) {
	var req listSharedlinkRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListSharedlinkParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sharedlinks, err := server.store.ListSharedlink(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sharedlinks)
}

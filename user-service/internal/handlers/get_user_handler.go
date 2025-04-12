package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// nolint:unused
type GetUserHandler struct {
	db *sqlx.DB
}

func (h *GetUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// nolint:unused
type CreateUserHandler struct {
	db *sqlx.DB
}

func (h *CreateUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

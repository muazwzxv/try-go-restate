package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/db/repository"
)

// nolint:unused
type GetUserHandler struct {
	db *sqlx.DB
}

func (h *GetUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := repository.GetUserByUUID(ctx, "uuid", h.db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.JSON(http.StatusNotFound, nil) // return dem response body
				return
			}
			ctx.JSON(http.StatusInternalServerError, nil) // return dem response body
			return
		}

		ctx.JSON(http.StatusOK, ResponseBody(user))
	}
}

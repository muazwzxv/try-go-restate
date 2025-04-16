package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/db/repository"
)

// nolint:unused
type GetUserHandler struct {
	db *sqlx.DB
}

var (
	ErrUserNotFound = ErrorDetail{
		HttpStatusCode: http.StatusNotFound,
		Message:        "USER_NOT_FOUND",
	}
	ErrUserxxx = ErrorDetail{}
)

func (h *GetUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := repository.GetUserByUUID(ctx, "uuid", h.db)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("error querying user, err: %v", err))
			if errors.Is(err, sql.ErrNoRows) {
				ctx.JSON(ErrUserNotFound.HttpStatusCode, ErrorResponse(ErrUserNotFound))
				return
			}
			ctx.JSON(ErrInternalError.HttpStatusCode, ErrorResponse(ErrInternalError))
			return
		}

		ctx.JSON(http.StatusOK, ResponseBody(user))
	}
}

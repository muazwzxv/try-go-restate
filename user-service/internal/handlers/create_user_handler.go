package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/internal/application/workflow"
)

// nolint:unused
type CreateUserHandler struct {
	db              *sqlx.DB
	WorkflowInvoker workflow.Invoker
}

type CreateUserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     string `json:"age"`
	Address string `json:"address"`
}

func (h *CreateUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody CreateUserRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.JSON(ErrBadRequest.HttpStatusCode, ErrorResponse(ErrBadRequest))
			return
		}

		// TODO: API call to restate server to invoke handler

		// TODO: kept track of invocation with the idempotency key
	}
}

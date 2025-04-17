package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/internal/application/workflow"
	"github.com/muazwzxv/try-go-restate/user-service/internal/workers"
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

type CreateUserResponse struct {
	ReferenceID string `json:"referenceID"`
	Email       string `json:"email"`
	Status      string `json:"status"`
}

func (h *CreateUserHandler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idempotencyKey := ctx.GetHeader(string(HeaderIdempotencyV1))
		var reqBody CreateUserRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.JSON(ErrBadRequest.HttpStatusCode, ErrorResponse(ErrBadRequest))
			return
		}

		workflowReq := &workers.CreateUserRequest{
			Name:  reqBody.Name,
			Email: reqBody.Email,
		}

		jsonByte := workflowReq.ToJSON()
		if jsonByte != nil {
			slog.Error("error creating request payload")
			ctx.JSON(ErrBadRequest.HttpStatusCode, ErrorResponse(ErrBadRequest))
			return
		}

		resp, err := h.WorkflowInvoker.InvokeWorkflow(context.Background(), &workflow.InvocationParams{
			Workflow:       workflow.WorkflowCreateUser,
			Payload:        jsonByte,
			IdempotencyKey: idempotencyKey,
		})
		if err != nil {
			slog.Error(fmt.Sprintf("error triggering workflow, err: %v", err))
			ctx.JSON(ErrInternalError.HttpStatusCode, ErrorResponse(ErrInternalError))
		}

		ctx.JSON(http.StatusOK, ResponseBody(&CreateUserResponse{
			ReferenceID: resp.InvocationID,
			Email:       reqBody.Email,
			Status:      resp.Status,
		}))
	}
}

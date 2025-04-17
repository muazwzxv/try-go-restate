package handlers

import (
	"net/http"

	app "github.com/muazwzxv/try-go-restate/user-service/internal/application"
)

var (
	ErrBadRequest = ErrorDetail{
		HttpStatusCode: http.StatusBadRequest,
		Message:        "BAD_REQUEST",
	}

	ErrInternalError = ErrorDetail{
		HttpStatusCode: http.StatusInternalServerError,
		Message:        "INTERNAL_SERVER_ERROR",
	}
)

type Handlers struct {
	GetUser    GetUserHandler
	CreateUser CreateUserHandler
}

func NewHandlers(app *app.Application) *Handlers {
	getUserHandler := GetUserHandler{
		db: app.DB,
	}

	createUserHandler := CreateUserHandler{
		db:              app.DB,
		WorkflowInvoker: nil, // TODO: register invoker workflow
	}

	return &Handlers{
		GetUser:    getUserHandler,
		CreateUser: createUserHandler,
	}
}

func ResponseBody(data any) map[string]any {
	return map[string]any{
		"data": data,
	}
}

func ErrorResponse(errDetail any) map[string]any {
	return map[string]any{
		"error": errDetail,
	}
}

type ErrorDetail struct {
	HttpStatusCode int    `json:"code"`
	Message        string `json:"message"`
}

type HeaderKeys string

var HeaderIdempotencyV1 HeaderKeys = "x-idempotency-key-v1"

package handlers

import app "github.com/muazwzxv/try-go-restate/user-service/internal/application"

type Handlers struct {
	GetUser    GetUserHandler
	CreateUser CreateUserHandler
}

func NewHandlers(app *app.Application) *Handlers {
	getUserHandler := GetUserHandler{
		db: app.DB,
	}

	createUserHandler := CreateUserHandler{
		db: app.DB,
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

func ErrorResponse(err error) map[string]any {
	return map[string]any{
		"error": err.Error(),
	}
}

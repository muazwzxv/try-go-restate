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

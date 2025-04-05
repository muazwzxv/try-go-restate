package handlers

import (
	"fmt"
	"log/slog"

	"github.com/muazwzxv/try-go-restate/user-service/internal/entities"
	restate "github.com/restatedev/sdk-go"
)

type CreateUser struct{}

//nolint:unused
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (CreateUser) CreateUser(ctx restate.Context, req *createUserRequest) (*entities.UserEntity, error) {
	slog.Info(fmt.Sprintf("Payload coming in, %v", req))
	return &entities.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

package workers

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/db/repository"
	"github.com/muazwzxv/try-go-restate/user-service/internal/entities"
	restate "github.com/restatedev/sdk-go"
)

type UserServiceWorkflows struct {
	DB *sqlx.DB
}

//nolint:unused
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *CreateUserRequest) ToJSON() []byte {
	if c == nil {
		return nil
	}

	jsonData, err := json.Marshal(c)
	if err != nil {
		slog.Error("Failed to marshal JSON", "err", err.Error())
		return nil
	}

	return jsonData
}

func (w UserServiceWorkflows) ExecuteCreateUserWorkflow(ctx restate.Context, req *CreateUserRequest) (*entities.UserEntity, error) {
	// TODO: ensure no existing email
	_, err := repository.GetUserByEmail(ctx, req.Email, w.DB)
	if err != nil {
		// TODO: handle error the restate way
		return nil, nil
	}

	// TODO: create user

	slog.Info(fmt.Sprintf("Payload coming in, %v", req))
	return &entities.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

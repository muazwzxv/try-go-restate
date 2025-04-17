package workers

import (
	"database/sql"
	"encoding/json"
	"errors"
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

var (
	ErrUserExist = errors.New("USER_EXIST_ERROR")
	ErrDatabase  = errors.New("DATABASE_ERROR")
)

func (w UserServiceWorkflows) ExecuteCreateUserWorkflow(ctx restate.Context, req *CreateUserRequest) (*entities.UserEntity, error) {
	user, err := repository.GetUserByEmail(ctx, req.Email, w.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, ErrDatabase
	}

	if err == nil && user.Status == "ACTIVE" {
		return nil, ErrUserExist
	}

	slog.Info(fmt.Sprintf("Payload coming in, %v", req))

	// TODO: create user entry in the database
	return &entities.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

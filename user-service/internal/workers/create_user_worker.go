package workers

import (
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
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (w UserServiceWorkflows) ExecuteCreateUserWorkflow(ctx restate.Context, req *createUserRequest) (*entities.UserEntity, error) {
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

package workers

import (
	"fmt"
	"log/slog"

	"github.com/muazwzxv/try-go-restate/user-service/internal/entities"
	restate "github.com/restatedev/sdk-go"
)

type CreateUserWorkflow struct{}

//nolint:unused
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (CreateUserWorkflow) ExecuteCreateUserWorkflow(ctx restate.Context, req *createUserRequest) (*entities.UserEntity, error) {
	// RPC call to other handlers
	// resp, err := restate.Object[any](ctx, "service-name", "key", "method").Request(restate.Void{})
	// if err != nil {
	// 	// Handle error
	// }

	slog.Info(fmt.Sprintf("Payload coming in, %v", req))
	return &entities.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/internal/entities/dto"
	"golang.org/x/net/context"
)

type UserModel struct {
	ID     int64  `db:"id"`
	UUID   string `db:"uuid"`
	Name   string `db:"name"`
	Email  string `db:"email"`
	Status string `db:"status"`

	CreatedAt sql.NullTime `db:"created_at"`
	CreatedBy string       `db:"created_by"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	UpdatedBy string       `db:"updated_by"`
}

var (
	getUserByUUID = `
    SELECT 
      uuid, name, email, status, created_at,
      created_by, updated_at, updated_by
    FROM user 
    WHERE uuid = ?`

	getUserByEmail = `
    SELECT 
      uuid, name, email, status, created_at,
      created_by, updated_at, updated_by
    FROM user 
    WHERE email = ?`

	// nolint:unused
	createUser = `
    INSERT INTO user
      (uuid, name, email, status, created_by, updated_by)
    VALUES
      (?, ?, ?, ?, ?, ?)`
)

func GetUserByUUID(ctx context.Context, uuid string, db *sqlx.DB) (*UserModel, error) {
	user := &UserModel{}
	row := db.QueryRowxContext(ctx, getUserByUUID, uuid)

	if row.Err() != nil {
		return nil, row.Err()
	}

	if err := row.StructScan(user); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string, db *sqlx.DB) (*UserModel, error) {
	user := &UserModel{}
	row := db.QueryRowxContext(ctx, getUserByEmail, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	if err := row.StructScan(user); err != nil {
		return nil, err
	}

	return user, nil
}

func InsertUser(ctx context.Context, req *dto.CreateUserDto, db *sqlx.DB) error {
	_, err := db.ExecContext(ctx, createUser,
		req.UUID,
		req.Name,
		req.Email,
		req.Status,
		req.CreatedBy,
		req.UpdatedBy,
	)
	if err != nil {
		return err
	}
	return nil
}

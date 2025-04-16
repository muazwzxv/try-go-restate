package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
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
    INSERT
      uuid, name, email, status, created_at, updated_by
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

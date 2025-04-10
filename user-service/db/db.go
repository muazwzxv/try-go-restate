package db

import (

	// need init to run
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	dsn := "root:password@(localhost:3306)/restate_demo"

	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}


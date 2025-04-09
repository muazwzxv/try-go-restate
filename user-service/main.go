package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/db"
	"github.com/muazwzxv/try-go-restate/user-service/internal/workers"
	restate "github.com/restatedev/sdk-go"
	"github.com/restatedev/sdk-go/server"
)

type application struct {
	db     *sqlx.DB
	server *server.Restate
}

func main() {
	app := setup()
	if err := app.server.Start(context.Background(), ":9090"); err != nil {
		slog.Error("application exited unexpectedly,", "err:", err.Error())
		os.Exit(1)
	}
}

func setup() *application {
	db, err := db.NewDB()
	if err != nil {
		slog.Error(fmt.Sprintf("error connecting db, err: %v", err))
		os.Exit(1)
	}

	// TODO: create Gin mux

	// create restate server for workflows
	server := server.NewRestate().
		Bind(restate.Reflect(workers.CreateUserWorkflow{}))

	return &application{
		db:     db,
		server: server,
	}
}

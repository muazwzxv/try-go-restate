package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/muazwzxv/try-go-restate/user-service/internal/application"
)

func main() {
	app := application.Setup()

	go func() {
		if err := app.Server.ListenAndServe(); err != nil {
			slog.Error("API server exited unexpectedly,", "err:", err.Error())
			os.Exit(1)
		}
	}()

	if err := app.Worker.Start(context.Background(), ":9090"); err != nil {
		slog.Error("worker exited unexpectedly,", "err:", err.Error())
		os.Exit(1)
	}
}

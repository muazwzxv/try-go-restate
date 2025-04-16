package application

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/muazwzxv/try-go-restate/user-service/db"
	"github.com/muazwzxv/try-go-restate/user-service/internal/workers"
	restate "github.com/restatedev/sdk-go"
	"github.com/restatedev/sdk-go/server"
)

type Application struct {
	DB     *sqlx.DB
	Worker *server.Restate
	Mux    *gin.Engine
	Server *http.Server
}

func Setup() *Application {
	db, err := db.NewDB()
	if err != nil {
		slog.Error(fmt.Sprintf("error connecting db, err: %v", err))
		os.Exit(1)
	}

	mux := NewMux()
	appServer := NewServer()

	// create restate server for workflows
	restateServer := server.NewRestate().
		Bind(restate.Reflect(workers.UserServiceWorkflows{
			DB: db,
		}))

	return &Application{
		DB:     db,
		Worker: restateServer,
		Mux:    mux,
		Server: appServer,
	}
}

func NewServer() *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", 3000),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Duration(7) * time.Second,
		WriteTimeout: time.Duration(7) * time.Second,
	}
	return server
}

func NewMux() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	// if cfg.Envrionment == "production" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(
		requestid.New(
			requestid.WithCustomHeaderStrKey("x-request-id"),
		),
	)

	return r
}

package app

import (
	"backend/internal/user"
	"backend/router"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Router *router.API
}

func New(db *pgxpool.Pool) *App {
	// USER module
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	api := router.NewAPI(userHandler)

	return &App{
		Router: api,
	}
}

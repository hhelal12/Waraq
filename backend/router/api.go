package router

import "backend/internal/user"

type API struct {
	UserHandler *user.Handler
}

func NewAPI(userHandler *user.Handler) *API {
	return &API{
		UserHandler: userHandler,
	}
}
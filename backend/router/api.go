package router

import (
    "backend/internal/user"
)

type API struct {
    UserHandler   *user.Handler
}
package routes

import (
	"myveda/handlers"
	"myveda/pkg/middleware"
	"myveda/pkg/mysql"
	"myveda/repositories"

	"github.com/labstack/echo/v4"
)

func RouteAuth(e *echo.Group) {
	repositories := repositories.AuthRepo(mysql.DB)
	h := handlers.HandlerAuth(repositories)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("check-auth", middleware.Auth(h.CheckAuth))
}

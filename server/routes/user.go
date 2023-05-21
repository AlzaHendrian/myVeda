package routes

import (
	"myveda/handlers"
	"myveda/pkg/mysql"
	"myveda/repositories"

	"github.com/labstack/echo/v4"
)

func RouteUsers(e *echo.Group) {
	Repo := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUsers(Repo)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}

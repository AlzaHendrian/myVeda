package routes

import (
	"myveda/handlers"
	"myveda/pkg/middleware"
	"myveda/pkg/mysql"
	"myveda/repositories"

	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Group) {
	repository := repositories.ProductRepo(mysql.DB)
	h := handlers.ProductRepo(repository)

	e.GET("/product", h.FindProduct)
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
}

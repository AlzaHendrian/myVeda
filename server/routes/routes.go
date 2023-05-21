package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	RouteUsers(e)
	RouteProduct(e)
	RouteAuth(e)
	RouteTransaction(e)
}

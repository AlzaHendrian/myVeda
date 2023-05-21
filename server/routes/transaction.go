package routes

import (
	"myveda/handlers"
	"myveda/pkg/middleware"
	"myveda/pkg/mysql"
	"myveda/repositories"

	"github.com/labstack/echo/v4"
)

func RouteTransaction(e *echo.Group) {
	repoTransactions := repositories.TransactionRepo(mysql.DB)
	h := handlers.TransactionRepo(repoTransactions)

	e.GET("/transaction-by-login", middleware.Auth(h.FindTransactionByLogin))
	e.GET("/transactions", middleware.Auth(h.FindTransactions))
}

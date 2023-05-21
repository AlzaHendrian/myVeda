package handlers

import (
	dto "myveda/dto/result"
	"myveda/repositories"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	transaction repositories.RepoTransaction
}

func TransactionRepo(transaction repositories.RepoTransaction) *handlerTransaction {
	return &handlerTransaction{transaction}
}

func (h *handlerTransaction) FindTransactionByLogin(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	transactions, err := h.transaction.FindTransactionByLogin(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.transaction.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

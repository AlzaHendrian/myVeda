package transaction

type CreateTransactionRequest struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Status    string `json:"status" vaildate:"required"`
	Total     int    `json:"total" validate:"required"`
}

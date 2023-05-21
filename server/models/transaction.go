package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Total     int       `json:"total"`
	Status    string    `json:"status"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	ProductID int       `json:"product_id"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

type Favorite struct {
	ID        int     `json:"id" gorm:"primary_key:auto_increment"`
	Product   Product `json:"product"`
	ProductID int     `json:"product_id"`
	User      User    `json:"user"`
	UserID    int     `json:"user_id"`
}

type FavoriteResponse struct {
	ID      int     `json:"id"`
	Product Product `json:"product"`
}

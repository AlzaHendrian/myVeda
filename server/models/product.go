package models

type Product struct {
	ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
	Title     string              `json:"title" gorm:"type: varchar(255)"`
	Thumbnail string              `json:"thumbnail" gorm:"type: varchar(255)"`
	Image     string              `json:"image" gorm:"type: varchar(255)"`
	Price     int                 `json:"price"`
	Size      string              `json:"size" gorm:"type: varchar(255)"`
	Style     string              `json:"style" gorm:"type: varchar(255)"`
	Material  string              `json:"material" gorm:"type: varchar(255)"`
	Color     string              `json:"color" gorm:"type: varchar (255)"`
	Condition string              `json:"condition" gorm:"type: varchar(255)"`
	UserID    int                 `json:"user_id"`
	User      UserProductResponse `json:"user"`
}

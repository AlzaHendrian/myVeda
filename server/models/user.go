package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key: auto_increment"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Fullname string `json:"fullname" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
}

type UserProductResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
}

func (UserProductResponse) TableName() string {
	return "users"
}

package usersdto

type CreateUser struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UpdateUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

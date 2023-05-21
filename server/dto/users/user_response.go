package usersdto

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type UserDeleteResponse struct {
	ID int `json:"id"`
}

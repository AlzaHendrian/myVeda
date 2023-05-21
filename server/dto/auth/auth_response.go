package authdto

type RespLogin struct {
	ID    int    `json:"id"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(255)" json:"role"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type RespRegister struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

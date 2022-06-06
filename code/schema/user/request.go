package schema

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Login
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type UserOrderCreate struct {
	Name string `json:"name" binding:"required"`
}

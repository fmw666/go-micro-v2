package schema

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReq struct {
	LoginReq
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type UserOrderCreateReq struct {
	Name string `json:"name" binding:"required"`
}

type UserResp struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

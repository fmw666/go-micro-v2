package schema

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	LoginReq
	PasswordConfirm string `json:"password_confirm"`
}

type UserResp struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

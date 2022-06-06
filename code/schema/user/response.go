package schema

type UserDetail struct {
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserDetailWithToken struct {
	User  UserDetail `json:"user"`
	Token string     `json:"token"`
}

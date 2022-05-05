package schema

type OrderCreateReq struct {
	Name string `json:"name"`
}

type OrderResp struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

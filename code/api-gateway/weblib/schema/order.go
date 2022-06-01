package schema

type OrderCreateReq struct {
	Name   string `json:"name"`
	UserID uint32 `json:"user_id"`
}

type OrderResp struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	UserID    uint32 `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

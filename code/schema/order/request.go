package schema

type OrderCreate struct {
	Name   string `json:"name"`
	UserID uint32 `json:"user_id"`
}

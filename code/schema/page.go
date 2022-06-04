package schema

type PageInfo struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
	Total  uint32 `json:"total"`
}

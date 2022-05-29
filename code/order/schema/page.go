package schema

type PageInfoResp struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
	Total  uint32 `json:"total"`
}

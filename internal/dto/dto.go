package dto

type BaseResponse struct {
	ID        uint  `json:"id" form:"id"`
	CreatedAt int64 `json:"createdAt" form:"createdAt"`
	UpdatedAt int64 `json:"updatedAt" form:"updatedAt"`
}

type PaginationData struct {
	Page  uint `form:"page"`
	Limit uint `form:"limit"`
}

type FindByFieldData struct {
	Key   string `form:"key"`
	Value any    `form:"value"`
}

package model

type Request_Pagination struct {
	OrderBy string `json:"orderBy" example:""`
	Order   string `json:"order" example:"ASC"`

	Page int64 `example:"1" json:"page"`
	Size int64 `example:"11" json:"size"`
}

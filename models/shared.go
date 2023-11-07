package models

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page  int32 `json:"page"`
	Size  int32 `json:"size"`
	Total int32 `json:"total"`
}

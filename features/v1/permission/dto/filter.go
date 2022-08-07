package dto

type Filter struct {
	Name        *string `json:"name" form:"name"`
	Description *string `json:"description" form:"description"`
}

package dto

type PermissionResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
type ListPermissionResponse []PermissionResponse

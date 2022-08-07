package dto

type CreatePermissionRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

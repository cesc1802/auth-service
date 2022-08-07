package dto

type UpdatePermissionRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

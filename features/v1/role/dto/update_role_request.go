package dto

type UpdateRoleRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

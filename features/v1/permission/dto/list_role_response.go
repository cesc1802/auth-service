package dto

type RoleResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
type ListRoleResponse []RoleResponse

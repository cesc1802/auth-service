package dto

type CreateRolePermissionRequest struct {
	RoleID      uint `json:"role_id"`
	Permissions []struct {
		ID uint `json:"id"`
	} `json:"permissions"`
}

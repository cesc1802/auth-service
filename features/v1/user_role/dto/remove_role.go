package dto

type RemoveRolesRequest struct {
	UserID uint `json:"user_id"`
	Roles  []struct {
		ID uint `json:"id"`
	} `json:"roles"`
}

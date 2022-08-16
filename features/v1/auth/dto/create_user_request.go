package dto

type CreateUserRequest struct {
	LoginID   string `json:"login_id"`
	Password  string `json:"password"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

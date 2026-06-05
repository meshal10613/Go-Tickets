package dto

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserTokenResponse struct {
	Token string       `json:"token,omitempty"`
	User  UserResponse `json:"user"`
}

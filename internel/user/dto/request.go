package dto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" query:"name" validate:"required"`
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required,password"`
}

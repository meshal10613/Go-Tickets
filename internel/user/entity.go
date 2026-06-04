package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name" query:"name" validate:"required" gorm:"type:varchar(100) not null"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email" gorm:"type:varchar(100); uniqueIndex; not null"`
	// Password string `json:"password" form:"password" query:"password" validate:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789,containsany=!@#$%^&*()_+-=[]{}|;:,.<>?/"`
	Password string `json:"password" form:"password" query:"password" validate:"required,password" gorm:"type:varchar(255) not null"`
}
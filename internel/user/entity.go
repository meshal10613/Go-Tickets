package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name" query:"name" gorm:"type:varchar(100) not null"`
	Email string `json:"email" form:"email" query:"email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Password string `json:"password" form:"password" query:"password" gorm:"type:varchar(255) not null"`
}
package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("User with this email already exists")

type Repository interface {
	Register(user *User) error
	GetByEmail(email string) (*User, error)
	// CheckPassword(password string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Register(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}

		return result.Error
	}

	return nil
}

func (r *repository) GetByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// func (r *userRepository) CheckPassword(password string) error {
// 	//? Implement password checking logic here
// 	return nil
// }

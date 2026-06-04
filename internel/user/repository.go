package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("User with this email already exists")

type UserRepository interface {
	CreateUser(user *User) error
	// GetUserByEmail(email string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}

		return result.Error
	}

	return nil
}

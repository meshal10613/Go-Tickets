package user

import "go-tickets/internel/user/dto"

type service struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	var err error
	user := User{
		Name:     req.Name,
		Email:    req.Email,
	}

	//? Hash password before saving to database
	err = user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}

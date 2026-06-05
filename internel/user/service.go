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
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.CreateUser(&user)
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

package user

import (
	"errors"
	"go-tickets/internel/auth"
	"go-tickets/internel/user/dto"
)

var ErrorInvalidCredentials = errors.New("Invalid email or password")

type service struct {
	repo       UserRepository
	jwtService auth.JwtService
}

func NewUserService(repo UserRepository, jwtService auth.JwtService) *service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *service) RegisterUser(req *dto.RegisterUserRequest) (*dto.UserTokenResponse, error) {
	var err error
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	//? Hash password before saving to database
	err = user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.RegisterUser(&user)
	if err != nil {
		return nil, err
	}

	// //? generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	response := dto.UserTokenResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return &response, nil
}

func (s *service) LoginUser(req *dto.LoginUserRequest) (*dto.UserTokenResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = user.checkPassword(req.Password)
	if user == nil || err != nil {
		return nil, ErrorInvalidCredentials
	}

	//? generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	response := dto.UserTokenResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return &response, nil
}

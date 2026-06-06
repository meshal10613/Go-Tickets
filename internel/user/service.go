package user

import (
	"errors"
	"go-tickets/internel/auth"
	"go-tickets/internel/user/dto"
)

var ErrorInvalidCredentials = errors.New("Invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JwtService
}

func NewService(repo Repository, jwtService auth.JwtService) *service {
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

	err = s.repo.Register(&user)
	if err != nil {
		return nil, err
	}

	// //? generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(token), nil
}

func (s *service) LoginUser(req *dto.LoginUserRequest) (*dto.UserTokenResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
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

	return user.ToResponse(token), nil
}

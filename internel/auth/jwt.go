package auth

import (
	"fmt"
	"go-tickets/internel/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(userID uint, name, email string) (string, error)
	ValidateToken(tokenString string) (*JwtClaims, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtService(secretKey string) JwtService {
	cfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %v", err))
	}

	if secretKey == "" {
		secretKey = cfg.JwtSecretKey
	}

	tokenDuration, err := ParseDuration()
	if err != nil {
		tokenDuration = 24 * time.Hour
	}

	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

func (js *jwtService) GenerateToken(userID uint, name, email string) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		Name:   name,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-tickets",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (js *jwtService) ValidateToken(tokenString string) (*JwtClaims, error) {
	// token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{},
	// 	func(token *jwt.Token) (interface{}, error) {
	// 		return []byte(js.secretKey), nil
	// 	},
	// )
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(js.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*JwtClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to extract claims")
}

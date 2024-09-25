package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthServiceInterface interface {
	Login(username, password string) (string, error)
	Register(username, password, email string) error
	GenerateToken(username, roles, email string) (string, error)
	ParseToken(token string) (*jwt.Token, error)
}

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthService struct {
	SecretKey *string
}

func NewAuthService(SecretKey *string) *AuthService {
	return &AuthService{SecretKey: SecretKey}
}

func (s *AuthService) Login(username, password string) (string, error) {
	return username, nil
}

func (s *AuthService) Register(username, password, email string) error {
	return nil
}

func (s *AuthService) GenerateToken(username, roles, email string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Role:     roles,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(*s.SecretKey))
	return tokenString, err
}
func (s *AuthService) ParseToken(tokenString string) (*jwt.Token, error) {
	var result CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &result, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.SecretKey, nil
	})
	return token, err
}

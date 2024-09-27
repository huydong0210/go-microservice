package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	http2 "go-microservices/cmd/auth/internal/api/handler/http"
	"go-microservices/internal/api/request"
	error2 "go-microservices/internal/error"
	"go-microservices/internal/helper"
	"time"
)

type AuthServiceInterface interface {
	Login(username, password string) (*string, error)
	Register(request *request.UserCreationRequest) error
	GenerateToken(id uint, username, roles, email string) (string, error)
	ParseToken(token string) (*jwt.Token, error)
}

type CustomClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthService struct {
	SecretKey   *string
	httpHandler *http2.HttpHandler
}

func NewAuthService(SecretKey *string, httpHandler *http2.HttpHandler) *AuthService {
	return &AuthService{SecretKey: SecretKey, httpHandler: httpHandler}
}

func (s *AuthService) Login(username, password string) (*string, error) {
	userInfo, err := s.httpHandler.GetUserInfoByUsername(username)
	if err != nil {
		return nil, err
	}
	if !helper.CheckPasswordHash(password, userInfo.HashPassword) {
		return nil, error2.NewAppError("Password incorect")
	}
	token, err := s.GenerateToken(userInfo.Id, userInfo.Username, userInfo.Roles, userInfo.Email)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *AuthService) Register(request *request.UserCreationRequest) error {
	return s.httpHandler.CreateUser(request)
}

func (s *AuthService) GenerateToken(id uint, username, roles, email string) (string, error) {
	claims := CustomClaims{
		Id:       id,
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

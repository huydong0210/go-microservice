package handler

import (
	"github.com/gin-gonic/gin"
	request2 "go-microservices/cmd/auth/internal/api/request"
	"go-microservices/cmd/auth/internal/service"
	request3 "go-microservices/internal/api/request"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request request2.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authService.Login(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (h *AuthHandler) Register(c *gin.Context) {
	var request request3.UserCreationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.authService.Register(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "register successfully"})
}
func (h *AuthHandler) ParseToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	splitToken := strings.Split(tokenString, " ")

	tokenString = splitToken[1]
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}
	token, err := h.authService.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims, oke := token.Claims.(*service.CustomClaims)
	if !oke {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": claims.Username,
		"email":    claims.Email,
		"role":     claims.Role,
	})

}

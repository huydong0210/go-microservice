package handler

import (
	"github.com/gin-gonic/gin"
	request2 "go-microservices/cmd/auth/internal/api/request"
	"go-microservices/cmd/auth/internal/service"
	"net/http"
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
}

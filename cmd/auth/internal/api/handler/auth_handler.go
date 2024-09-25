package handler

import (
	"github.com/gin-gonic/gin"
	http2 "go-microservices/cmd/auth/internal/api/handler/http"
	request2 "go-microservices/cmd/auth/internal/api/request"
	"go-microservices/cmd/auth/internal/service"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
	httpHandler *http2.HttpHandler
}

func NewAuthHandler(authService service.AuthServiceInterface, httpHandler *http2.HttpHandler) *AuthHandler {
	return &AuthHandler{authService: authService, httpHandler: httpHandler}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request request2.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userInfo, err := h.httpHandler.GetUserInfoByUsername(request.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	token, err := h.authService.GenerateToken(userInfo.Username, userInfo.Roles, userInfo.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

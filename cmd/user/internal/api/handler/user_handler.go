package handler

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/user/internal/service"
	request2 "go-microservices/internal/api/request"
	"go-microservices/internal/api/response"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService service.UserServiceInterface
	roleService service.RoleServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface, roleService service.RoleServiceInterface) *UserHandler {
	return &UserHandler{userService: userService, roleService: roleService}
}

func (h *UserHandler) FindUserLoginByUserName(c *gin.Context) {
	user, err := h.userService.FindUserByUserName(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	roles, err := h.roleService.FindAllRolesByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var roleNames string

	for _, role := range roles {
		roleNames += role.Name + " "
	}
	roleNames = strings.TrimSpace(roleNames)

	c.JSON(http.StatusOK, gin.H{
		"data": &response.UserInfoResponse{
			Username:     user.Username,
			HashPassword: user.Password,
			Roles:        roleNames,
			Email:        user.Email,
		},
	})
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var request request2.UserCreationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := h.userService.FindUserByUserName(request.Username)

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
		return
	}
	if err := h.userService.CreateUser(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

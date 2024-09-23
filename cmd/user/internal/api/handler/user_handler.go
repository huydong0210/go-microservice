package handler

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/user/internal/api/response"
	"go-microservices/cmd/user/internal/service"
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

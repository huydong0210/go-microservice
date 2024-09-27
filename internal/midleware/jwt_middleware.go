package midleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	http2 "go-microservices/internal/api/http"
	"net/http"
	"strings"
)

type UserPrincipal struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

const USER_PRICIPAL_CONTEXT_KEY = "USER_PRINCIPAL"

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}
		user, err := parseToken(tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		c.Set(USER_PRICIPAL_CONTEXT_KEY, user)

		c.Next()
	}
}

func parseToken(token string) (*UserPrincipal, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token
	request, err := http2.MakeRequest("http://localhost:8082/api/auth/token", http.MethodGet, nil, header)

	if err != nil {
		return nil, err
	}
	res, err := http2.DoRequest(request)
	if err != nil {
		return nil, err
	}
	var result UserPrincipal
	err = json.Unmarshal([]byte(*res), &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

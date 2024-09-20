package midleware

import (
	"github.com/gin-gonic/gin"
	"go-microservices/internal/service"
)

func WhiteIpMiddleware(service *service.WhiteIpService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}

}

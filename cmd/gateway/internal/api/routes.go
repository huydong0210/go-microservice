package api

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/gateway/internal/config"
	"net/http/httputil"
	"net/url"
)

func SetUpRoutes(router *gin.Engine, cfg *config.Config) {
	router.Any("/api/auth/*any", func(context *gin.Context) {
		proxy := createReverseProxy(createUri(cfg.Host, cfg.AuthServiceAddress))
		proxy.ServeHTTP(context.Writer, context.Request)
	})

	router.Any("/api/user/*any", func(context *gin.Context) {
		proxy := createReverseProxy(createUri(cfg.Host, cfg.UserServiceAddress))
		proxy.ServeHTTP(context.Writer, context.Request)
	})

	router.Any("/api/todo-item/*any", func(context *gin.Context) {
		proxy := createReverseProxy(createUri(cfg.Host, cfg.TodoServiceAddress))
		proxy.ServeHTTP(context.Writer, context.Request)
	})
}

func createUri(host, address string) string {
	return "http://" + host + address
}

func createReverseProxy(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(url)
}

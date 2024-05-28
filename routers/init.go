package routers

import (
	"flow/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	initRouter(r)

	return r
}

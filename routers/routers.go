package routers

import (
	"github.com/gin-gonic/gin"

	"flow/apis/handler"
	"flow/apis/test"
)

func initRouter(r *gin.Engine) {
	initTestRouter(r)
	initUnAuthRouter(r)
}

func initTestRouter(r *gin.Engine) {
	v1 := r.Group("/api/test")

	v1.GET("/hello", test.Hello)
}

func initUnAuthRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/getCaptcha", handler.GenerateCaptcha)

}

package routers

import (
	"github.com/gin-gonic/gin"

	"flow/apis/controllers"
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
	v1.GET("/getCaptcha", controllers.GenerateCaptcha)
	v1.POST("/auth/register", controllers.Register)
}

func initApplicationRouter(v1 *gin.RouterGroup, r *gin.Engine) {
	app := v1.Group("/app")
	{
		app.GET("detail")
		app.POST("/register", controllers.Register)
	}
}

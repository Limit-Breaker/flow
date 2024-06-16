package controllers

import (
	"flow/apis/services"
	"flow/common/request"
	"flow/common/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	reg := &request.Register{}
	if err := c.ShouldBind(reg); err != nil {
		response.ValidateFail(c, err.Error())
		return
	}

	if err, user := services.UserService.Register(reg); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

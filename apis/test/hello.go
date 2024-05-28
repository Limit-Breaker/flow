package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Hello say hello
func Hello(c *gin.Context) {
	name := c.Query("name")
	response := struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
		Code   int    `json:"code"`
	}{
		Status: "",
		Msg:    "hello " + name,
		Code:   0,
	}
	c.JSON(http.StatusOK, response)
}

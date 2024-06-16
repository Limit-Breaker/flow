package response

import (
	"flow/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type Response struct {
//	// 代码
//	Code int `json:"code" example:"10000"`
//	// 数据集
//	Data interface{} `json:"data"`
//	// 消息
//	Msg string `json:"msg"`
//}
//
//func (r *Response) ResponseOK() *Response {
//	r.Code = 0
//	return r
//}
//
//func (r *Response) ResponseErr(code int) *Response {
//	r.Code = code
//	return r
//}

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Data    interface{} `json:"data"`    // 数据
	Message string      `json:"message"` // 信息
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		data,
		"ok",
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error common.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, common.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, common.Errors.BusinessError.ErrorCode, msg)
}

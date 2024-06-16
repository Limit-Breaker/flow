package controllers

import (
	"flow/common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

func GenerateCaptcha(c *gin.Context) {
	id, b64s, err := driverDigitFunc()
	if err != nil {
		fmt.Println("generate captcha failed: ", err)
		response.BusinessFail(c, err.Error())
	}
	data := gin.H{
		"id":      id,
		"captcha": b64s,
	}

	response.Success(c, data)
}

func driverDigitFunc() (id, b64s string, err error) {
	// configJsonBody json request body.
	type configJsonBody struct {
		Id            string
		CaptchaType   string
		VerifyValue   string
		DriverAudio   *base64Captcha.DriverAudio
		DriverString  *base64Captcha.DriverString
		DriverChinese *base64Captcha.DriverChinese
		DriverMath    *base64Captcha.DriverMath
		DriverDigit   *base64Captcha.DriverDigit
	}
	var store = base64Captcha.DefaultMemStore

	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	driver := e.DriverDigit
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err = captcha.Generate()
	return
}

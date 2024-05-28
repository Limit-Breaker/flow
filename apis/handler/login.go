package handler

import (
	"flow/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func GenerateCaptcha(c *gin.Context) {
	id, b64s, err := DriverDigitFunc()
	res := common.Response{}
	if err != nil {
		fmt.Println("generate captcha failed: ", err)
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ResponseErr(10001))
	}

	data := gin.H{
		"id":      id,
		"captcha": b64s,
	}
	res.Data = data
	c.JSON(http.StatusOK, res.ResponseOK())
}

func DriverDigitFunc() (id, b64s string, err error) {
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

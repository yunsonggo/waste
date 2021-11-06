package common

import (
	"github.com/mojocn/base64Captcha"
)

type GenerateStrCaptchaBody struct {
	Id   string
	B64s string
}

type VerifyCaptchaBody struct {
	Id          string
	VerifyValue string
}

var store = base64Captcha.DefaultMemStore

func GenerateCaptchaStr() (CaptchaCode *GenerateStrCaptchaBody, err error) {
	var driverConfig = base64Captcha.DriverDigit{
		40,
		120,
		4,
		0.4,
		30,
	}
	var DriverDigit = base64Captcha.NewDriverDigit(driverConfig.Height, driverConfig.Width, driverConfig.Length, driverConfig.MaxSkew, driverConfig.DotCount)
	var driver base64Captcha.Driver
	driver = DriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	var Code GenerateStrCaptchaBody
	Code.Id = id
	Code.B64s = b64s
	return &Code, nil
}

func VerifyCaptchaStr(code VerifyCaptchaBody) (res bool) {
	return store.Verify(code.Id, code.VerifyValue, true)
}

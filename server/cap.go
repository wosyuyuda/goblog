package server

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

//生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          50,
		Width:           120,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          3,
		Source:          "1234567890",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

//验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	return store.Verify(id, capt, false)
	/* if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	} */
	//return
}

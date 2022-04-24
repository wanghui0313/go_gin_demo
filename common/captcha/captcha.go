package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func Captcha(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		driverString := base64Captcha.DriverString{
			Height:          60,
			Width:           200,
			NoiseCount:      0,
			ShowLineOptions: 2 | 4,
			Length:          4,
			Source:          "123456789abcdefghijkmnopqrstuvwxyz",
			BgColor: &color.RGBA{
				R: 5,
				G: 50,
				B: 100,
				A: 150,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		}
		var driver base64Captcha.Driver = driverString.ConvertFonts()
		c := base64Captcha.NewCaptcha(driver, store)
		id, b64s, err := c.Generate()
		body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
		if err != nil {
			body = map[string]interface{}{"code": 0, "msg": err.Error()}
		}
		context.JSON(200, body)
	}
}

func VerifyCaptcha(r *gin.Engine) gin.HandlerFunc {
	return func(context *gin.Context) {
		captcha := context.PostForm("cptcha")
		captcha_id := context.PostForm("cptcha_id")
		body := map[string]interface{}{"code": 0, "msg": "failed"}
		if store.Verify(captcha_id, captcha, true) {
			body = map[string]interface{}{"code": 1, "msg": "ok"}
		}
		context.JSON(200, body)
	}
}

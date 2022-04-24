package router

import (
	"github.com/gin-gonic/gin"
	"server/common/captcha"
)

func captchaRouter(r *gin.Engine) {
	r.GET("/captcha", captcha.Captcha(r))
	r.POST("/verifyCaptcha", captcha.VerifyCaptcha(r))
}

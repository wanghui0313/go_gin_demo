package router

import (
	"github.com/gin-gonic/gin"
	"server/controller/user"
	"server/middler"
	"server/middler/jwt"
)

func userRouter(r *gin.Engine) {
	route := r.Group("user")
	{
		route.GET("index", middler.LoginMiddle, user.Index(r))
		route.POST("login", user.Login(r))
		route.GET("getUser", jwt.JwtAuth, user.GetUserInfo)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"server/common"
)

var router *gin.Engine

func init() {
	router = gin.New()
	//Miss路由
	router.NoRoute(func(context *gin.Context) {
		context.String(404, "没有此路由")
		context.Abort()
		return
	})
	//注册函数给模板使用
	router.SetFuncMap(template.FuncMap{
		"sum": common.Sum,
	})
	//指定模板文件
	router.LoadHTMLGlob("view/**/*")
	//全局中间件
	//router.Use(middler.AllMiddle)
}

func InitRouter() {
	indexRouter(router)
	userRouter(router)
	goodsRouter(router)
	captchaRouter(router) //验证码
	staticRouter(router)  //静态文件
	uploadRouter(router)
	router.Run(":80")
}

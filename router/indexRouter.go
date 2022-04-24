package router

import (
	"github.com/gin-gonic/gin"
	"server/controller/index"
)

func indexRouter(r *gin.Engine) {
	r.GET("/", index.Index(r))
	route := r.Group("index")
	{
		route.GET("index", index.Index(r))
	}
}

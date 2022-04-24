package router

import (
	"github.com/gin-gonic/gin"
	"server/controller/goods"
)

func goodsRouter(r *gin.Engine) {
	route := r.Group("goods")
	{
		route.GET("index", goods.Index(r))
	}
}

package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/index", gin.H{
			"title": "goodstitle",
		})
	}
}

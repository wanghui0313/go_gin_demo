package index

import (
	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "index/index", gin.H{})
	}
}

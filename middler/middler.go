package middler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllMiddle(c *gin.Context) {
	//c.String(http.StatusOK, "all middle\n")
	fmt.Println("all middle\n")
}

func AllMiddle2(c *gin.Context) {
	c.String(http.StatusOK, "all middle2222\n")
}

func LoginMiddle(c *gin.Context) {
	c.Set("name", "wanghui123")
	//c.Redirect(302, "https://www.baidu.com")
}

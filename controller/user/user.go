package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"server/common"
	"server/common/token"
	admin2 "server/model/admin"
)

type User struct {
	Uid      int
	Name     string `form:"name"  binding:"required,mobile"`
	Password string `form:"pass"  binding:"required,min=6,max=12"`
}

func Index(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Request.URL.Path = "/goodsList"
		//r.HandleContext(c)
		c.HTML(http.StatusOK, "user/index", gin.H{
			"title": "wanghuihuiwang",
		})
		//c.String(200, c.GetString("name"))
	}
}

func Login(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mobile", common.CheckMobile)
		}
		var u User
		err := c.ShouldBind(&u)
		if err != nil {
			c.String(200, err.Error())
			return
		}

		if u.Password != "123456" || u.Name != "17601237968" {
			c.JSON(http.StatusOK, gin.H{"code": 401, "msg": "login failed"})
			return
		}
		u.Uid = 888
		token := token.GetToken(u.Uid)
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "login success", "token": token})
	}
}

func GetUserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	admin := admin2.GetAdmin()
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "GetUserInfo success", "claims": claims, "admin": admin})
}

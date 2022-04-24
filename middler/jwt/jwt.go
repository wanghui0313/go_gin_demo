package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/common/token"
)

func JwtAuth(c *gin.Context) {
	tokenString := c.Request.Header.Get("token")
	if tokenString == "" {
		c.JSON(404, gin.H{"code": 404, "msg": "no token"})
		c.Abort()
		return
	}
	Claims := token.UserClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &Claims, func(t *jwt.Token) (interface{}, error) {
		return token.JWTKEY, nil
	})
	if err != nil {
		ve, _ := err.(*jwt.ValidationError)
		if ve.Errors == jwt.ValidationErrorExpired {
			c.JSON(404, gin.H{"code": 404, "msg": "token expired"})
		} else {
			c.JSON(404, gin.H{"code": 404, "token": "token invalid"})
		}
		c.Abort()
		return
	}
	c.Set("claims", Claims)
}

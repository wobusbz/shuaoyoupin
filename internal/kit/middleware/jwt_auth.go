package middleware

import (
	"net/http"
	"shuaoyoupin/internal/kit/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizations := strings.Split(c.GetHeader("Authorization"), " ")
		if len(authorizations) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未登录或非法访问"})
			c.Abort()
			return
		}
		if authorizations[0] != "banner" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "非法授权头部"})
			c.Abort()
			return
		}
		authToken, err  := jwt.NewAuthToken().DecodeToken(authorizations[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "效验失败"})
			c.Abort()
			return
		}
		c.Set("auth", authToken.UserInter)
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,AccessToken,"+
			"X_Requested_With,Accept, Origin, Host, "+
			"x-user-id,Token,Connection, Accept-Encoding, "+
			"Accept-Language,DNT, X-CustomHeader, Keep-Alive, "+
			"User-Agent, X-Requested-With, If-Modified-Since, "+
			"XFILENAME,XFILECATEGORY,XFILESIZE"+
			"Cache-Control, Content-Type, Pragma,token,openid,"+
			"opentoken,Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		ctx.Header("Access-Control-Max-Age", "172800")
		ctx.Header("Access-Control-Allow-Credentials", "false")

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,x-requested-with,XFILENAME,XFILECATEGORY,XFILESIZE,AccessToken,X-CSRF-Token, x-user-id,Authorization, Token,x-token")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, x-requested-with,XFILENAME,XFILECATEGORY,XFILESIZE,Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}

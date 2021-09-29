package router

import (
	"shuaoyoupin/internal/router/api"

	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	engine := gin.Default()
	var router = engine.Group("")
	(&api.ApiRouter{
		
	}).RegisterApiRouter(router)
	return engine
}

package api

import (
	v1Api "shuaoyoupin/api/v1/api"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	apiUser *v1Api.V1Api
}

func (ur *ApiRouter) RegisterApiRouter(r *gin.RouterGroup) {
	r.POST("userCreate", ur.apiUser.Register)
}

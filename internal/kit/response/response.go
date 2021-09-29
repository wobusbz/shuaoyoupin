package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseParamFailed(ctx *gin.Context, format string, args ...interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(format, args...)})
	return
}

func ResponseOk(ctx *gin.Context, obj interface{}, format string, args ...interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": obj,
		"message": fmt.Sprintf(format, args...), 
	})
	return
}

func ResponseFailed(ctx *gin.Context, format string, args ...interface{}) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf(format, args...)})
	return
}

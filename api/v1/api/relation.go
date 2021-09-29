package v1Api

import (
	"github.com/gin-gonic/gin"
	"shuaoyoupin/internal/kit/response"
	"shuaoyoupin/internal/model"
)

type (
	InvitationCodeResponse struct {
		InvitationCode int `json:"invitation_code,omitempty"`
	}
)

func (v1 *V1Api) RelationGetInvitationCode(ctx *gin.Context) {
	claims, ok := ctx.Get("userInfo")
	if !ok {
		response.ResponseParamFailed(ctx, "Token 不存在")
		return
	}
	userInfo := claims.(*model.ApiUser)
	resultUser, err := v1.srvRelation.RelationGet(map[string]interface{}{"user_id": userInfo.ID})
	if err != nil {
		response.ResponseFailed(ctx, "用户不存在")
		return
	}
	response.ResponseOk(ctx, &InvitationCodeResponse{InvitationCode: resultUser.InvitationCode}, "Ok")
}

package v1Api

import (
	"shuaoyoupin/internal/kit/jwt"
	"shuaoyoupin/internal/kit/password"
	"shuaoyoupin/internal/kit/random"
	"shuaoyoupin/internal/kit/response"
	"shuaoyoupin/internal/kit/verify"
	"shuaoyoupin/internal/model"

	"github.com/gin-gonic/gin"
)

type (
	ApiUserCreateRequest struct {
		InvitationCode int    `json:"invitation_code,omitempty"`
		Mobile         string `json:"mobile,omitempty"`
		PassWorld      string `json:"pass_world,omitempty"`
	}
	ApiUserLoginRequest struct {
		Mobile    string `json:"mobile,omitempty"`
		PassWorld string `json:"pass_world,omitempty"`
	}
	ApiUserLoginResponse struct {
		Token string `json:"token,omitempty"`
	}
)

func (v1 *V1Api) Login(ctx *gin.Context) {
	var req = &ApiUserCreateRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.ResponseParamFailed(ctx, "参数错误: %v", err)
		return
	}
	resultUser, err := v1.srvUser.ApiUserGet(map[string]interface{}{"mobile": req.Mobile})
	if err != nil {
		response.ResponseFailed(ctx, "用户不存在")
		return
	}
	if err := password.VerifyPassword(resultUser.PassWord, req.PassWorld); err != nil {
		response.ResponseFailed(ctx, "用户账户或密码错误")
		return
	}
	token, err := jwt.NewAuthToken().EncodeToken(resultUser)
	if err != nil {
		response.ResponseFailed(ctx, "Token 加密失败")
		return
	}
	response.ResponseOk(ctx, &ApiUserLoginResponse{Token: token}, "Ok")
}

func (v1 *V1Api) Register(ctx *gin.Context) {
	var req = &ApiUserCreateRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.ResponseParamFailed(ctx, "参数错误: %v", err)
		return
	}
	if !verify.VerifyMobile(req.Mobile) {
		response.ResponseParamFailed(ctx, "手机号码错误")
		return
	}
	if req.InvitationCode <= 0 {
		response.ResponseParamFailed(ctx, "邀请码不能为空")
		return
	}
	if _, err := v1.srvUser.ApiUserGet(map[string]interface{}{"mobile": req.Mobile}); err == nil {
		response.ResponseFailed(ctx, "用户已存在")
		return
	}
TO:
	invitationCode := random.InstanceRandom().AssignNInt(6)
	resultParent, err := v1.srvRelation.RelationGet(map[string]interface{}{"code": invitationCode})
	if err != nil {
		goto TO
	}
	pass, err := password.EncryptionPassword(req.PassWorld)
	if err != nil {
		response.ResponseFailed(ctx, "用户加密失败")
		return
	}
	var resultUser = &model.ApiUser{Mobile: req.Mobile, PassWord: pass}

	if err := v1.srvUser.ApiUserCreate(resultUser); err != nil {
		response.ResponseFailed(ctx, "注册失败")
		return
	}
	resultParent.InvitationCode = req.InvitationCode
	resultParent.ID = resultUser.ID
	if err := v1.srvRelation.RelationCreate(resultParent); err != nil {
		response.ResponseFailed(ctx, "注册关系链失败")
		return
	}
	response.ResponseOk(ctx, nil, "Ok")
}

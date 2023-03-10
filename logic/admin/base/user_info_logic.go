package base

import (
	"go-flow-admin/pkg/jwt"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/admin/base"
)

// UserInfo 获取用户信息
func UserInfo(ctx *svc.ServiceContext) (resp base.UserInfoResponse, err error) {
	user, _ := ctx.GinContext.Get("userInfo")
	claims := user.(*jwt.Claims)
	userModel := query.AdminUserModel
	resp = userModel.GetUserInfo(claims.Id)
	if claims.RoleId == 1 {
		resp.Authkeys = "*"
	}
	return
}

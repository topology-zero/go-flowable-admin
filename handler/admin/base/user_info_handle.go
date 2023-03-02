package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// UserInfoHandle 获取用户信息
func UserInfoHandle(c *gin.Context) {
	res, err := base.UserInfo(svc.NewServiceContext(c))
	response.HandleResponse(c, res, err)
}

package login

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/login"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// LogoutHandle 退出登录
func LogoutHandle(c *gin.Context) {
	err := login.Logout(svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

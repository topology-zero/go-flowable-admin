package login

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/login"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// CodeHandle 获取验证码
func CodeHandle(c *gin.Context) {
	resp, err := login.Code(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

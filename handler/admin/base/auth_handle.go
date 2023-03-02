package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// AuthHandle 获取所有权限
func AuthHandle(c *gin.Context) {
	resp, err := base.Auth(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

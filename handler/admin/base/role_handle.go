package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// RoleHandle 获取所有角色
func RoleHandle(c *gin.Context) {
	resp, err := base.Role(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// AdminUsersHandle 获取所有用户
func AdminUsersHandle(c *gin.Context) {
	resp, err := base.AdminUsers(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

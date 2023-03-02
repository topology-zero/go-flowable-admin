package role

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/role"
	"go-flow-admin/svc"
	roleType "go-flow-admin/types/admin/role"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑角色
func EditHandle(c *gin.Context) {
	var req roleType.RoleEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Edit(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

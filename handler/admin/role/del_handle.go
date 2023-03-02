package role

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/role"
	"go-flow-admin/svc"
	roleType "go-flow-admin/types/admin/role"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除角色
func DelHandle(c *gin.Context) {
	var req roleType.RoleDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

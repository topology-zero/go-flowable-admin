package role

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/role"
	"go-flow-admin/svc"
	roleType "go-flow-admin/types/admin/role"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加角色
func AddHandle(c *gin.Context) {
	var req roleType.RoleAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := role.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

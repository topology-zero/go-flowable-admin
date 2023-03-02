package role

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/role"
	"go-flow-admin/svc"
	roleType "go-flow-admin/types/admin/role"

	"github.com/gin-gonic/gin"
)

// ListHandle 角色列表
func ListHandle(c *gin.Context) {
	var req roleType.RoleListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	resp, err := role.List(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

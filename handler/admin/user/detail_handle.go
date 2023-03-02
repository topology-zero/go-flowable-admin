package user

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/user"
	"go-flow-admin/svc"
	userType "go-flow-admin/types/admin/user"

	"github.com/gin-gonic/gin"
)

// DetailHandle 用户详情
func DetailHandle(c *gin.Context) {
	var req userType.UserDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := user.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

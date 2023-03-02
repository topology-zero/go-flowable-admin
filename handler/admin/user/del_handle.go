package user

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/user"
	"go-flow-admin/svc"
	userType "go-flow-admin/types/admin/user"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除用户
func DelHandle(c *gin.Context) {
	var req userType.UserDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := user.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

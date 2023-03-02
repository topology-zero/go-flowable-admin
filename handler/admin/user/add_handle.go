package user

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/user"
	"go-flow-admin/svc"
	userType "go-flow-admin/types/admin/user"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加用户
func AddHandle(c *gin.Context) {
	var req userType.UserAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := user.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

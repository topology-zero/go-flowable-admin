package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/base"
	"go-flow-admin/svc"
	baseType "go-flow-admin/types/admin/base"

	"github.com/gin-gonic/gin"
)

// ChangeSelfPwdHandle 修改自己的密码
func ChangeSelfPwdHandle(c *gin.Context) {
	var req baseType.ChangeSelfPwdRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := base.ChangeSelfPwd(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

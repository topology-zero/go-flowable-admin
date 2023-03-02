package auth

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/auth"
	"go-flow-admin/svc"
	authType "go-flow-admin/types/admin/auth"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除权限
func DelHandle(c *gin.Context) {
	var req authType.AuthDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := auth.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

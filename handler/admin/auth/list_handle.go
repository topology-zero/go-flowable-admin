package auth

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/auth"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 权限列表
func ListHandle(c *gin.Context) {
	resp, err := auth.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

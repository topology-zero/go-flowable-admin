package generate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/generate"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// ListHandle 表结构列表
func ListHandle(c *gin.Context) {
	resp, err := generate.List(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

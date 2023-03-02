package generate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/admin/generate"
	"go-flow-admin/svc"
	generateType "go-flow-admin/types/admin/generate"

	"github.com/gin-gonic/gin"
)

// GenerateHandle 生成前端文件
func GenerateHandle(c *gin.Context) {
	var req generateType.GenerateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := generate.Generate(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

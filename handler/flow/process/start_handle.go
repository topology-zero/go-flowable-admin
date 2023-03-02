package process

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process"
	"go-flow-admin/svc"
	processType "go-flow-admin/types/flow/process"

	"github.com/gin-gonic/gin"
)

// StartHandle 启动流程
func StartHandle(c *gin.Context) {
	var req processType.StartProcessRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := process.Start(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

package process

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process"
	"go-flow-admin/svc"
	processType "go-flow-admin/types/flow/process"

	"github.com/gin-gonic/gin"
)

// DetailHandle 流程详情
func DetailHandle(c *gin.Context) {
	var req processType.DetailProcessRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := process.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

package instance

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process/instance"
	"go-flow-admin/svc"
	instanceType "go-flow-admin/types/flow/process/instance"

	"github.com/gin-gonic/gin"
)

// DetailHandle 流程详情
func DetailHandle(c *gin.Context) {
	var req instanceType.ProcessInstanceDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := instance.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

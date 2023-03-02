package process

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process"
	"go-flow-admin/svc"
	processType "go-flow-admin/types/flow/process"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加流程
func AddHandle(c *gin.Context) {
	var req processType.AddProcessRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := process.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

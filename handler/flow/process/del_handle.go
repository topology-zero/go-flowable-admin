package process

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process"
	"go-flow-admin/svc"
	processType "go-flow-admin/types/flow/process"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除流程
func DelHandle(c *gin.Context) {
	var req processType.DeleteProcessRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := process.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

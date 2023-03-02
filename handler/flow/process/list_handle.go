package process

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process"
	"go-flow-admin/svc"
	processType "go-flow-admin/types/flow/process"

	"github.com/gin-gonic/gin"
)

// ListHandle 流程列表
func ListHandle(c *gin.Context) {
	var req processType.ListProcessRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	resp, err := process.List(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

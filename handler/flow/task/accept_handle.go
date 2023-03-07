package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// AcceptHandle 审批同意
func AcceptHandle(c *gin.Context) {
	var req taskType.TaskAcceptRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.Accept(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

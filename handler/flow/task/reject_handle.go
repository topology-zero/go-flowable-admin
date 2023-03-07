package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// RejectHandle 审批驳回
func RejectHandle(c *gin.Context) {
	var req taskType.TaskRejectRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.Reject(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

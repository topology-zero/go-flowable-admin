package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// DetailHandle 任务详情
func DetailHandle(c *gin.Context) {
	var req taskType.TaskDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := task.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

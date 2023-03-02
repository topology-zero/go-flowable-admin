package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// CompleteHandle 完成任务
func CompleteHandle(c *gin.Context) {
	var req taskType.CompleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.Complete(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

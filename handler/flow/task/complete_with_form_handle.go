package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// CompleteWithFormHandle 提交表单完成任务
func CompleteWithFormHandle(c *gin.Context) {
	var req taskType.CompleteWithFormRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.CompleteWithForm(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

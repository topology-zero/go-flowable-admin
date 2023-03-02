package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// DelAttachmentHandle 删除任务附件
func DelAttachmentHandle(c *gin.Context) {
	var req taskType.DelAttachmentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.DelAttachment(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

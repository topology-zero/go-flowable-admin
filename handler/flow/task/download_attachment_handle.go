package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// DownloadAttachmentHandle 下载任务附件
func DownloadAttachmentHandle(c *gin.Context) {
	var req taskType.DownloadAttachmentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	data, err := task.DownloadAttachment(&req, svc.NewServiceContext(c))
	if err != nil {
		response.HandleResponse(c, nil, err)
	}
	c.Data(200, "application/octet-stream", data)
}

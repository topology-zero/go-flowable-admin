package task

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/task"
	"go-flow-admin/svc"
	taskType "go-flow-admin/types/flow/task"

	"github.com/gin-gonic/gin"
)

// DelCommentHandle 删除备注
func DelCommentHandle(c *gin.Context) {
	var req taskType.DelCommentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := task.DelComment(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// FlowcateHandle 流程分类列表
func FlowcateHandle(c *gin.Context) {
	resp, err := base.Flowcate(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/base"
	"go-flow-admin/svc"

	"github.com/gin-gonic/gin"
)

// FlowformHandle 流程表单列表
func FlowformHandle(c *gin.Context) {
	resp, err := base.Flowform(svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

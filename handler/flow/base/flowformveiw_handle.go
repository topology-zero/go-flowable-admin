package base

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/base"
	"go-flow-admin/svc"
	baseType "go-flow-admin/types/flow/base"

	"github.com/gin-gonic/gin"
)

// FlowformveiwHandle 流程表单预览
func FlowformveiwHandle(c *gin.Context) {
	var req baseType.FlowFormViewRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := base.Flowformveiw(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

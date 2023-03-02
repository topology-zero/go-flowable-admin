package form

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/form"
	"go-flow-admin/svc"
	formType "go-flow-admin/types/flow/form"

	"github.com/gin-gonic/gin"
)

// DetailHandle 流程外置表单详情
func DetailHandle(c *gin.Context) {
	var req formType.FlowFormDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := form.Detail(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

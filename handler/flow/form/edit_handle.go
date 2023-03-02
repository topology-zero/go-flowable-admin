package form

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/form"
	"go-flow-admin/svc"
	formType "go-flow-admin/types/flow/form"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑流程外置表单
func EditHandle(c *gin.Context) {
	var req formType.FlowFormEditRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := form.Edit(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

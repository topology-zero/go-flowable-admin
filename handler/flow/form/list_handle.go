package form

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/form"
	"go-flow-admin/svc"
	formType "go-flow-admin/types/flow/form"

	"github.com/gin-gonic/gin"
)

// ListHandle 流程外置表单列表
func ListHandle(c *gin.Context) {
	var req formType.FlowFormListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := form.List(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

package form

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/form"
	"go-flow-admin/svc"
	formType "go-flow-admin/types/flow/form"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加流程外置表单
func AddHandle(c *gin.Context) {
	var req formType.FlowFormAddRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := form.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

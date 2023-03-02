package form

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/form"
	"go-flow-admin/svc"
	formType "go-flow-admin/types/flow/form"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除流程外置表单
func DelHandle(c *gin.Context) {
	var req formType.FlowFormDeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := form.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

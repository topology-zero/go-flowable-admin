package instance

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/process/instance"
	"go-flow-admin/svc"
	instanceType "go-flow-admin/types/flow/process/instance"

	"github.com/gin-gonic/gin"
)

// AllHandle 全部流程
func AllHandle(c *gin.Context) {
	var req instanceType.ProcessInstanceRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := instance.All(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

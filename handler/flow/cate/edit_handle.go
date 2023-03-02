package cate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/cate"
	"go-flow-admin/svc"
	cateType "go-flow-admin/types/flow/cate"

	"github.com/gin-gonic/gin"
)

// EditHandle 编辑流程分类
func EditHandle(c *gin.Context) {
	var req cateType.EditFlowCateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := cate.Edit(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

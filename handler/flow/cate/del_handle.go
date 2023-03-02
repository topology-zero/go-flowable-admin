package cate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/cate"
	"go-flow-admin/svc"
	cateType "go-flow-admin/types/flow/cate"

	"github.com/gin-gonic/gin"
)

// DelHandle 删除流程分类
func DelHandle(c *gin.Context) {
	var req cateType.DeleteFlowCateRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := cate.Del(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

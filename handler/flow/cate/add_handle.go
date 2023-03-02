package cate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/cate"
	"go-flow-admin/svc"
	cateType "go-flow-admin/types/flow/cate"

	"github.com/gin-gonic/gin"
)

// AddHandle 添加流程分类
func AddHandle(c *gin.Context) {
	var req cateType.AddFlowCateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	err := cate.Add(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, nil, err)
}

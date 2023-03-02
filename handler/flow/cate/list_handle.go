package cate

import (
	"go-flow-admin/internal/response"
	"go-flow-admin/logic/flow/cate"
	"go-flow-admin/svc"
	cateType "go-flow-admin/types/flow/cate"

	"github.com/gin-gonic/gin"
)

// ListHandle 流程分类列表
func ListHandle(c *gin.Context) {
	var req cateType.FlowCateListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 1
	}

	resp, err := cate.List(&req, svc.NewServiceContext(c))
	response.HandleResponse(c, resp, err)
}

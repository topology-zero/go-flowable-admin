package cate

import (
	"go-flow-admin/pkg/util"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/cate"
)

// List 流程分类列表
func List(req *cate.FlowCateListRequest, ctx *svc.ServiceContext) (resp cate.FlowCateListResponse, err error) {
	flowCateModel := query.FlowCateModel
	data, count, _ := flowCateModel.FindByPage((req.Page-1)*req.PageSize, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = int(count)
	_ = util.CopyValue(&resp.Data, &data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}

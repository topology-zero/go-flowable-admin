package form

import (
	"github.com/jinzhu/copier"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// List 流程外置表单列表
func List(req *form.FlowFormListRequest, ctx *svc.ServiceContext) (resp form.FlowFormListResponse, err error) {
	formModel := query.FlowFormModel
	list, count, _ := formModel.FindByPage((req.Page-1)*req.PageSize, req.PageSize)
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = int(count)
	copier.Copy(&resp.Data, &list)
	for i := range list {
		resp.Data[i].Id = list[i].ID
	}

	return
}

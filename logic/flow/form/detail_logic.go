package form

import (
	"github.com/jinzhu/copier"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Detail 流程外置表单详情
func Detail(req *form.FlowFormDetailRequest, ctx *svc.ServiceContext) (resp form.FlowFormDetailResponse, err error) {
	formModel := query.FlowFormModel
	first, _ := formModel.Where(formModel.ID.Eq(req.Id)).First()
	copier.Copy(&resp, &first)
	resp.Id = first.ID
	return
}

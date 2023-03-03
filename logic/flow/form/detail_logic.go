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
	first, _ := formModel.Where(formModel.Key.Eq(req.Key), formModel.Version.Eq(req.Version)).First()
	copier.Copy(&resp, &first)
	return
}

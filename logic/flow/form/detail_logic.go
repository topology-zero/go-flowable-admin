package form

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_definition"
)

// Detail 流程外置表单详情
func Detail(req *form.FlowFormDetailRequest, ctx *svc.ServiceContext) (resp form.FlowFormDetailResponse, err error) {
	model, err := form_definition.Model(req.Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	resp.Rule = model.Fields[0].Value.(string)
	resp.Option = model.Fields[1].Value.(string)
	return
}

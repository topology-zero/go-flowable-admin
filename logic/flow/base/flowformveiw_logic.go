package base

import (
	"github.com/MasterJoyHunan/flowablesdk/external_form/form_definition"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/base"
)

// Flowformveiw 流程表单预览
func Flowformveiw(req *base.FlowFormViewRequest, ctx *svc.ServiceContext) (resp base.FlowFormViewResponse, err error) {
	model, err := form_definition.Model(req.Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	resp.Rule = model.Fields[0].Value.(string)
	resp.Option = model.Fields[1].Value.(string)
	return
}

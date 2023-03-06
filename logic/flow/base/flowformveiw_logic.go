package base

import (
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_definition"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/base"
)

// Flowformveiw 流程表单预览
func Flowformveiw(req *base.FlowFormViewRequest, ctx *svc.ServiceContext) (resp base.FlowFormViewResponse, err error) {
	param := form_definition.ListRequest{}
	param.Size = 1
	param.Latest = true
	param.Key = req.Key
	list, _, err := form_definition.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
		return
	}
	if len(list) != 1 {
		ctx.Log.Errorf("获取表单错误 list = %+v", list)
		err = errors.New("获取表单错误")
		return
	}
	model, err := form_definition.Model(list[0].Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
		return
	}
	resp.Rule = model.Fields[0].Value.(string)
	resp.Option = model.Fields[1].Value.(string)
	return
}

package base

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_definition"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/base"
)

// Flowform 流程表单列表
func Flowform(ctx *svc.ServiceContext) (resp base.FlowFormResponse, err error) {
	param := form_definition.ListRequest{}
	param.Size = 1000
	param.Sort = "version"
	param.Order = "desc"
	param.Latest = true
	list, _, err := form_definition.List(param)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	copier.Copy(&resp.List, &list)
	return
}

package base

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/base"
)

// Flowformveiw 流程表单预览
func Flowformveiw(req *base.FlowFormViewRequest, ctx *svc.ServiceContext) (resp base.FlowFormViewResponse, err error) {
	formModel := query.FlowFormModel
	first, err := formModel.Unscoped().Where(formModel.Key.Eq(req.Key), formModel.Version.Eq(req.Version)).First()
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("系统错误")
		return
	}
	copier.Copy(&resp, &first)
	return
}

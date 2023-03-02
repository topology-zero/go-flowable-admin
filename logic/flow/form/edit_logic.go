package form

import (
	"github.com/pkg/errors"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Edit 编辑流程外置表单
func Edit(req *form.FlowFormEditRequest, ctx *svc.ServiceContext) error {
	formModel := query.FlowFormModel
	_, err := formModel.Where(formModel.ID.Eq(req.Id)).UpdateColumnSimple(
		formModel.Name.Value(req.Name),
		formModel.Rule.Value(req.Rule),
		formModel.Option.Value(req.Option),
	)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		return errors.New("系统错误")
	}
	return nil
}

package form

import (
	"github.com/pkg/errors"
	"go-flow-admin/model"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Add 添加流程外置表单
func Add(req *form.FlowFormAddRequest, ctx *svc.ServiceContext) error {
	formModel := query.FlowFormModel
	err := formModel.Create(&model.FlowFormModel{
		Name:   req.Name,
		Rule:   req.Rule,
		Option: req.Option,
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		return errors.New("系统错误")
	}
	return nil
}

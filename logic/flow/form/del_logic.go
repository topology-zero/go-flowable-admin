package form

import (
	"github.com/pkg/errors"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Del 删除流程外置表单
func Del(req *form.FlowFormDeleteRequest, ctx *svc.ServiceContext) error {
	// TODO 需要判断是否在使用
	formModel := query.FlowFormModel
	_, err := formModel.Where(formModel.ID.Eq(req.Id)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		return errors.New("系统错误")
	}
	return nil
}

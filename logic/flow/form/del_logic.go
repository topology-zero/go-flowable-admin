package form

import (
	"github.com/MasterJoyHunan/flowablesdk/external_form/form_definition"
	"github.com/MasterJoyHunan/flowablesdk/external_form/form_deployment"
	"github.com/pkg/errors"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Del 删除流程外置表单
func Del(req *form.FlowFormDeleteRequest, ctx *svc.ServiceContext) error {
	detail, err := form_definition.Detail(req.Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("未能获取数据")
	}

	err = form_deployment.Delete(detail.DeploymentId)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("删除失败")
	}

	formModel := query.FlowFormModel
	_, err = formModel.Where(formModel.Key.Eq(detail.Key), formModel.Version.Eq(detail.Version)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}

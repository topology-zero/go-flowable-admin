package form

import (
	"github.com/MasterJoyHunan/flowablesdk/external_form/form_deployment"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Del 删除流程外置表单
func Del(req *form.FlowFormDeleteRequest, ctx *svc.ServiceContext) error {
	err := form_deployment.Delete(req.Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("删除失败")
	}
	return err
}

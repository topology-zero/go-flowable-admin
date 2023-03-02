package process

import (
	"github.com/MasterJoyHunan/flowablesdk/deployment"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"
)

// Del 删除流程
func Del(req *process.DeleteProcessRequest, ctx *svc.ServiceContext) error {
	err := deployment.Delete(req.Id)
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("删除流程失败")
		return err
	}
	return nil
}

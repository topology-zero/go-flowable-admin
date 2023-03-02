package process

import (
	"github.com/MasterJoyHunan/flowablesdk/deployment"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"
)

// Add 添加流程
func Add(req *process.AddProcessRequest, ctx *svc.ServiceContext) (resp process.AddProcessResponse, err error) {
	dep, err := deployment.Create(deployment.CreateRequest{
		FileName: req.FileName + ".bpmn20.xml",
		Category: req.Category,
		Xml:      req.Xml,
	})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("添加失败")
		return
	}
	resp.Id = dep.Id
	return
}

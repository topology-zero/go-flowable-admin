package process

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/deployment"
)

// Add 添加流程
func Add(req *process.AddProcessRequest, ctx *svc.ServiceContext) (resp process.AddProcessResponse, err error) {
	dep, err := deployment.Create(deployment.CreateRequest{
		FileName: req.FileName + ".bpmn20.xml",
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

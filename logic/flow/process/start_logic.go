package process

import (
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/process"

	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/process_instance"
	"github.com/topology-zero/flowablesdk/variable"
)

// Start 启动流程
func Start(req *process.StartProcessRequest, ctx *svc.ServiceContext) (err error) {
	_, err = process_instance.Start(process_instance.StartProcessRequest{
		ProcessDefinitionId: req.Id,
		Variables: []variable.VariableRequest{
			{
				Name:  "process_instance_var",
				Type:  "integer",
				Value: 1,
			},
		},
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("启动流程失败")
		return err
	}
	return nil
}

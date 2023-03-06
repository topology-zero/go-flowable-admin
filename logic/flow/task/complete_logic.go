package task

import (
	"github.com/pkg/errors"
	processTask "github.com/topology-zero/flowablesdk/task"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// Complete 完成任务
func Complete(req *task.CompleteRequest, ctx *svc.ServiceContext) error {
	err := processTask.Action(req.Id, processTask.ActionRequest{
		Action: "complete",
		//Variables: []variables.VariableRequest{
		//	{
		//		Name:  "testVar",
		//		Type:  "integer",
		//		Value: 1,
		//	},
		//},
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("完成任务错误")
		return err
	}
	return nil
}

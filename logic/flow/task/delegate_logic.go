package task

import (
	"strconv"

	"github.com/pkg/errors"
	processTask "github.com/topology-zero/flowablesdk/task"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// Delegate 指派任务
func Delegate(req *task.DelegateRequest, ctx *svc.ServiceContext) error {
	err := processTask.Action(req.Id, processTask.ActionRequest{
		Action:   "delegate",
		Assignee: strconv.Itoa(req.ToUser),
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("指派任务错误")
		return err
	}

	return nil
}

package task

import (
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/form"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// CompleteWithForm 提交表单完成任务
func CompleteWithForm(req *task.CompleteWithFormRequest, ctx *svc.ServiceContext) error {
	var prop []form.Properties
	for _, v := range req.Properties {
		prop = append(prop, form.Properties{
			Id:    v.Id,
			Value: v.Value,
		}, form.Properties{
			Id:    req.Id + "|" + v.Id,
			Value: v.Value,
		})
	}

	_, err := form.SubmitForm(form.SubmitFromRequest{
		TaskId:     req.Id,
		Properties: prop,
	})

	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("提交任务失败")
	}
	return err
}

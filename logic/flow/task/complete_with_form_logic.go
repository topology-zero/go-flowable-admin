package task

import (
	"github.com/MasterJoyHunan/flowablesdk/form"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/task"
)

// CompleteWithForm 提交表单完成任务
func CompleteWithForm(req *task.CompleteWithFormRequest, ctx *svc.ServiceContext) error {

	var prop []form.Properties

	//for _, p := range req.Properties {
	//	prop = append(prop, form.Properties{
	//		Id:    p.Id,
	//		Value: p.Value,
	//	})
	//}

	copier.Copy(&prop, &req.Properties)

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

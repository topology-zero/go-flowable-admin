package form

import (
	"github.com/pkg/errors"
	"github.com/topology-zero/flowablesdk/external_form/form_deployment"
	formReq "github.com/topology-zero/flowablesdk/external_form/model"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Add 添加流程外置表单
func Add(req *form.FlowFormAddRequest, ctx *svc.ServiceContext) error {
	fields := []formReq.FormField{
		{Id: "rule", Value: req.Rule},
		{Id: "option", Value: req.Option},
	}
	_, err := form_deployment.Add(form_deployment.AddRequest{
		FileName: req.Name + ".form",
		Data: formReq.Model{
			Name:        req.Name,
			Description: req.Desc,
			Key:         req.Key,
			Fields:      fields,
		},
	})
	if err != nil {
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("添加流程外置表单错误")
	}
	return err
}

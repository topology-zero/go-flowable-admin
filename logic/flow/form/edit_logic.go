package form

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk/external_form/form_deployment"
	formReq "github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/pkg/errors"
	"go-flow-admin/dto"
	"go-flow-admin/model"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Edit 编辑流程外置表单
func Edit(req *form.FlowFormEditRequest, ctx *svc.ServiceContext) error {
	var formRules []dto.FormCreateRule
	_ = json.Unmarshal([]byte(req.Rule), &formRules)

	fields := make([]formReq.FormField, len(formRules))
	for i := range formRules {
		fields[i].Id = formRules[i].Field
		fields[i].Name = formRules[i].Title
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

	formModel := query.FlowFormModel
	flowForm := &model.FlowFormModel{
		Key:     req.Key,
		Version: req.Version + 1,
		Rule:    req.Rule,
		Option:  req.Option,
	}
	err = formModel.Create(flowForm)
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}

package form

import (
	"encoding/json"
	"strconv"

	"github.com/MasterJoyHunan/flowablesdk/external_form/form_deployment"
	formReq "github.com/MasterJoyHunan/flowablesdk/external_form/model"
	"github.com/pkg/errors"
	"go-flow-admin/dto"
	"go-flow-admin/model"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/form"
)

// Add 添加流程外置表单
func Add(req *form.FlowFormAddRequest, ctx *svc.ServiceContext) error {
	tx := query.Q.Begin()
	formModel := query.FlowFormModel
	flowForm := &model.FlowFormModel{
		Rule:    req.Rule,
		Option:  req.Option,
		Version: 1,
	}
	err := tx.FlowFormModel.Create(flowForm)
	if err != nil {
		tx.Rollback()
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		return errors.New("系统错误")
	}

	var formRules []dto.FormCreateRule
	_ = json.Unmarshal([]byte(req.Rule), &formRules)

	fields := make([]formReq.FormField, len(formRules))
	for i := range formRules {
		fields[i].Id = formRules[i].Field
		fields[i].Name = formRules[i].Title
	}

	_, err = form_deployment.Add(form_deployment.AddRequest{
		FileName: req.Name + ".form",
		Data: formReq.Model{
			Name:        req.Name,
			Description: req.Desc,
			Key:         strconv.Itoa(flowForm.ID),
			Fields:      fields,
		},
	})
	if err != nil {
		tx.Rollback()
		ctx.Log.Errorf("%+v", errors.WithStack(err))
		err = errors.New("添加流程外置表单错误")
	}
	_, err = tx.FlowFormModel.Where(formModel.ID.Eq(flowForm.ID)).UpdateSimple(
		formModel.Key.Value(strconv.Itoa(flowForm.ID)),
	)
	if err != nil {
		tx.Rollback()
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	tx.Commit()
	return err
}

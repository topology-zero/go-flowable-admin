package cate

import (
	"github.com/pkg/errors"
	"go-flow-admin/model"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/cate"
)

// Add 添加流程分类
func Add(req *cate.AddFlowCateRequest, ctx *svc.ServiceContext) error {
	flowCateModel := query.FlowCateModel
	err := flowCateModel.Create(&model.FlowCateModel{
		Name: req.Name,
	})
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}

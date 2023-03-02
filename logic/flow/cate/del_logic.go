package cate

import (
	"github.com/pkg/errors"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/cate"
)

// Del 删除流程分类
func Del(req *cate.DeleteFlowCateRequest, ctx *svc.ServiceContext) error {
	flowCateModel := query.FlowCateModel
	_, err := flowCateModel.Where(flowCateModel.ID.Eq(req.Id)).Delete()
	if err != nil {
		ctx.Log.Errorf("数据库异常：%+v", errors.WithStack(err))
		err = errors.New("系统错误")
	}
	return err
}

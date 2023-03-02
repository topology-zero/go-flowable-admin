package base

import (
	"github.com/jinzhu/copier"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/flow/base"
)

// Flowcate 流程分类列表
func Flowcate(ctx *svc.ServiceContext) (resp base.FlowcateResponse, err error) {
	flowCateModel := query.FlowCateModel
	list, _ := flowCateModel.Find()
	_ = copier.Copy(&resp.List, &list)
	return
}

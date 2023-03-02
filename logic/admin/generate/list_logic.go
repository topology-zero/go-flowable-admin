package generate

import (
	"go-flow-admin/model"
	"go-flow-admin/svc"
	"go-flow-admin/types/admin/generate"
)

// List 表结构列表
func List(ctx *svc.ServiceContext) (resp generate.GenerateListResponse, err error) {
	model.DB().Raw("show tables").Scan(&resp.Tables)
	return
}

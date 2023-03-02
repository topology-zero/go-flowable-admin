package base

import (
	"github.com/jinzhu/copier"
	"go-flow-admin/query"
	"go-flow-admin/svc"
	"go-flow-admin/types/admin/base"
)

// Role 获取所有角色
func Role(ctx *svc.ServiceContext) (resp base.BaseRoleResponse, err error) {
	roleModel := query.AdminRoleModel
	data, _ := roleModel.Find()
	copier.Copy(&resp.Data, data)
	for i := range data {
		resp.Data[i].Id = data[i].ID
	}
	return
}
